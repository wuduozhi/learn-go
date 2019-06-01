package files

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConfig() (password, path, sqlLogin string) {
	var xxx = map[string]string{}
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return
	}
	return xxx["password"], xxx["path"], xxx["sqlLogin"]
}

// const sqlLogin = "root:meng22@tcp(localhost:3306)/ldxy?charset=utf8"

// Ad is the pic to be shown, contains the URL, Expire time, and the Order of ad.
type Ad struct {
	MD5    string `json:"md5"`
	URL    string `json:"URL"`
	Start  string `json:"start"`
	Expire string `json:"expire"`
	Long   int64  `json:"long"`
	Order  int16  `json:"order"`
	Free   bool   `json:"free"`
	Onair  bool   `json:"onair"`
}

type ByOrder []Ad

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

func DbList(reserve bool) (str []byte, err error) {
	_, _, sqlLogin := GetConfig()
	db, err := sql.Open("mysql", sqlLogin)
	defer db.Close()
	if err != nil {
		return
	}
	rows, err := db.Query("select * from files")
	defer rows.Close()
	if err != nil {
		return
	}
	var ads []Ad
	startTime := ""
	expireTime := ""
	loc, _ := time.LoadLocation("Local")
	for rows.Next() {
		var ad Ad
		err = rows.Scan(&ad.MD5, &ad.URL, &startTime, &expireTime, &ad.Order, &ad.Long, &ad.Free)
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, loc)
		t2, _ := time.ParseInLocation("2006-01-02 15:04:05", expireTime, loc)
		if t1.After(time.Now()) {
			ad.Onair = false
			if reserve {
				continue
			}
		} else if t2.Before(time.Now()) {
			ad.Onair = false
			_, err = db.Exec("delete from files where MD5=?", ad.MD5)
			if err != nil {
				return
			}
			continue
		} else {
			ad.Onair = true
		}
		ad.URL = "img/" + ad.URL
		ad.Start = t1.Format("2006-01-02 15:04:05")
		ad.Expire = t2.Format("2006-01-02 15:04:05")
		ads = append(ads, ad)
	}

	sort.Sort(ByOrder(ads))

	if !reserve {
		for i, v := range ads {
			_, err = db.Exec("update files set `order`=? where `MD5`=?", i+1, v.MD5)
			if err != nil {
				return
			}
		}
	}

	str, err = json.Marshal(ads)
	return
}

func TvList() (str []byte, err error) {
	_, _, sqlLogin := GetConfig()
	type Tv struct {
		Id         string `json:"id"`
		LastOnline string `json:"last_online"`
		Name       string `json:"name"`
		Online     bool   `json:"online"`
	}
	db, err := sql.Open("mysql", sqlLogin)
	defer db.Close()
	if err != nil {
		return
	}
	rows, err := db.Query("select * from tvs")
	defer rows.Close()
	if err != nil {
		return
	}

	var tvs []Tv
	loc, _ := time.LoadLocation("Local")
	for rows.Next() {
		var tv Tv
		lastonline := ""
		err = rows.Scan(&tv.Id, &lastonline, &tv.Name)
		if err != nil {
			return
		}
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", lastonline, loc)
		tv.LastOnline = t1.Format("2006-01-02 15:04:05")
		tmp := t1.Add(time.Minute * 5)
		if tmp.Before(time.Now()) {
			tv.Online = false
		} else {
			tv.Online = true
		}
		tvs = append(tvs, tv)
	}
	str, err = json.Marshal(tvs)
	return
}

// MD5Sum does calculate and shortten the md5 string
func MD5Sum(f multipart.File) (str string, err error) {
	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		return
	}
	f.Seek(0, 0)
	str = fmt.Sprintf("%x", h.Sum(nil))
	tmpString := []byte(str)
	str = string(tmpString[0:8])
	return
}
