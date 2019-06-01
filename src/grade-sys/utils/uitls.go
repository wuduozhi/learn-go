package utils

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"grade-sys/pkg/logging"
	"grade-sys/models"
)



func GetGrade(xn int,xq int,stuid string,hdjwpass string,ptpass string,force int) (models.Resp,error) {
	// _url := "http://spider.qnxg.net/bks/grade?xn=%d&xq=%d&stuid=%s&password=%s&ptPassword=%s&force=%d"
	// url := fmt.Sprintf(_url,xn,xq,stuid,hdjwpass,ptpass,force)

	_url := "http://spider.qnxg.net/bks/grade?xn=%d&xq=%d&stuid=%s&password=%s&force=%d"
	url := fmt.Sprintf(_url,xn,xq,stuid,hdjwpass,force)

	logging.Info(url)
	resp, err := http.Get(url)
	checkErr(err)
	
	defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
	
	var result models.Resp
	
	err = json.Unmarshal(body, &result)

	if err != nil {
		logging.Info(url,err,string(body))
		return result,err
	}
	return result,err
}

func checkErr(err error){
	if err != nil {
		logging.Debug(err)
	}
}
