package api

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"ldxy/files"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Signup handles tv signup...
func Signup(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	password, _, sqlLogin := files.GetConfig()
	type Resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	if r.Method == "POST" {
		clientpass := r.PostFormValue("password")
		passbytes := []byte(password)
		chk := fmt.Sprintf("%x", md5.Sum(passbytes))
		if chk != clientpass {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		id := r.PostFormValue("id")
		name := r.PostFormValue("name")

		db, err := sql.Open("mysql", sqlLogin)
		defer db.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = db.Exec("insert into tvs (id,name) values(?,?)", id, name)
		fmt.Printf("%+v", err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			resp := Resp{
				Status:  "ok",
				Message: "Successfully added!"}
			str, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusOK)
			w.Write(str)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ListHandler sends the adv list.
func ListHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	if r.Method == "GET" {
		advlist, err := files.DbList(true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(advlist)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// BeatHandler handles the heart beat of the TVs
func BeatHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	_, _, sqlLogin := files.GetConfig()
	type beatResp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}

	err := r.ParseMultipartForm(100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.PostFormValue("id")
	db, err := sql.Open("mysql", sqlLogin)
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rows, err := db.Query("select id from tvs where id=?", id)
	defer rows.Close()
	if !rows.Next() {
		str, _ := json.Marshal(beatResp{
			Status: "Failed",
			Error:  "ID not found!"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(str)
		return
	}

	// update the last beat of a tv.
	t1 := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec("update tvs set last_online=? where id=?", t1, id)

	jsonStr, _ := json.Marshal(beatResp{
		Status: "ok",
		Error:  "null"})
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStr)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	password, path, sqlLogin := files.GetConfig()
	switch r.Method {
	case "GET":
		advlist, err := files.DbList(false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(advlist)
	case "POST":
		err := r.ParseMultipartForm(100000000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		startStr := r.PostFormValue("start")
		expireStr := r.PostFormValue("expire")
		start, _ := time.Parse("2006-01-02 15:04:05", startStr)
		expire, _ := time.Parse("2006-01-02 15:04:05", expireStr)
		free, _ := strconv.ParseInt(r.PostFormValue("free"), 10, 64)
		long, _ := strconv.ParseInt(r.PostFormValue("long"), 10, 64)
		order, _ := strconv.ParseInt(r.PostFormValue("order"), 10, 64)
		//get the *fileheaders
		//get a ref to the parsed multipart form
		m := r.MultipartForm
		filelist := m.File["uploadfile"]
		//for each fileheader, get a handle to the actual file
		if len(filelist) < 1 {
			http.Error(w, "You didn't upload any file!", http.StatusInternalServerError)
			return
		}

		file, err := filelist[0].Open()
		defer file.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		md5Str, err := files.MD5Sum(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := ""

		if len(strings.Split(filelist[0].Filename, ".")) < 2 {
			http.Error(w, "filename has no suffix!", http.StatusInternalServerError)
			return
		}
		filesplit := strings.Split(filelist[0].Filename, ".")
		suffix := filesplit[len(filesplit)-1]
		switch suffix {
		case "jpg", "jpeg", "png", "mp4":
			filename = md5Str + "." + suffix
		default:
			http.Error(w, "."+suffix+" file is not supported!", http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(path + filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db, err := sql.Open("mysql", sqlLogin)
		defer db.Close()
		_, err = db.Exec("insert into files values(?,?,?,?,?,?,?)", md5Str, filename, start, expire, order, long, free)
		if err != nil {
			http.Error(w, "File duplicated: "+err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{\"status\":\"ok\"}")
			return
		}
	case "DELETE":
		err := r.ParseMultipartForm(100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		md5Str := r.PostFormValue("md5")
		clientpass := r.PostFormValue("password")
		passbytes := []byte(password)
		chk := fmt.Sprintf("%x", md5.Sum(passbytes))
		if chk != clientpass {
			http.Error(w, "Password incorrect, please check!", http.StatusInternalServerError)
			return
		}

		db, err := sql.Open("mysql", sqlLogin)
		defer db.Close()
		rows, err := db.Query("select URL from files where MD5=?", md5Str)
		defer rows.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var url string
		for rows.Next() {
			rows.Scan(&url)
		}
		if url == "" {
			http.Error(w, "unExitedFile!", http.StatusInternalServerError)
			return
		}
		err = os.Remove(path + url)

		_, err = db.Exec("delete from files where MD5=?", md5Str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "{\"status\":\"ok\"}")
			return
		}
	case "PUT":
		type Ad struct {
			MD5    string `json:"md5"`
			URL    string `json:"URL"`
			Start  string `json:"start"`
			Expire string `json:"expire"`
			Long   int64  `json:"long"`
			Order  int16  `json:"order"`
			Free   bool   `json:"free"`
		}
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 10000))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data []Ad
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		db, err := sql.Open("mysql", sqlLogin)
		defer db.Close()
		for _, v := range data {
			startTime, _ := time.Parse("2006-01-02 15:04:05", v.Start)
			expireTime, _ := time.Parse("2006-01-02 15:04:05", v.Expire)
			_, err = db.Exec("update files set `start`=?,`expire`=?,`long`=?,`order`=?,`free`=? where `MD5`=?", startTime, expireTime, v.Long, v.Order, v.Free, v.MD5)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{\"status\":\"ok\"}")
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TvHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tvlist, err := files.TvList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(tvlist)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	password, _, _ := files.GetConfig()
	type resp struct {
		Status string `json:"status"`
		Time   string `json:"time"`
	}
	if r.Method == "POST" {
		clientpass := r.PostFormValue("password")
		passbytes := []byte(password)
		chk := fmt.Sprintf("%x", md5.Sum(passbytes))
		if chk == clientpass {
			w.WriteHeader(http.StatusOK)
			bytes, _ := json.Marshal(resp{
				Status: "ok",
				Time:   time.Now().Add(time.Hour * 1).Format("2006-01-02T15:04:05")})
			w.Write(bytes)
		} else {
			http.Error(w, "Password Incorrect!", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
