package main

import (
	"io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)



func main(){
	spider := HdjwSpider{
		Stuid:"201626010520",
		Password:"WudUozhI",
	}

	spider.login()
}


type HdjwSpider struct {
	Stuid string
	Password string
	Token string
}

func (spider *HdjwSpider) login(){
	url := "http://hdjw.hnu.edu.cn/secService/login"
	postdict := map[string]interface{}{
		"userCode":spider.Stuid,
		"password":spider.Password,
		"kaptcha":"testa",
		"userCodeType":"account",
	}

	bytesRepresentation,err := json.Marshal(postdict)
	fmt.Printf("%s \n", bytesRepresentation)
	client := &http.Client{}
	request ,err := http.NewRequest("POST",url,bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Add("app", "PCWEB")
	request.Header.Add("host", "hdjw.hnu.edu.cn")
	request.Header.Add("Referer", "http://hdjw.hnu.edu.cn/Njw2017/login.html")
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("KAPTCHA-KEY-GENERATOR-REDIS","securityKaptchaRedisServiceAdapter")
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	request.Header.Add("Accept","application/json, text/plain, */*")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	
	body,_ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	log.Println(result["data"])
	data := result["data"].(map[string]interface{})
	log.Println(data["token"])
	spider.Token = data["token"].(string)
}

