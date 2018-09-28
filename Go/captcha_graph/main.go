package main

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
)

type CaptchaInfo struct {
    Base64			string		`json:"base64"`
	CaptchaID 		string		`json:"captchaID"`
}


func showFormHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/captcha" {
        http.NotFound(w, r)
        return
    }
    var result CaptchaInfo
    result.Base64, result.CaptchaID = captcha.New()//在此处写入存储
    bytes, _ := json.Marshal(result)
    w.WriteHeader(http.StatusOK)
    w.Write(bytes)
}

func imageVerify(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/captcha/verify" {
		http.NotFound(w, r)
		return
	}
	var jsonReq map[string]string
	jsonBytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(jsonBytes, &jsonReq)
	captchaID:=jsonReq["captchaID"]
	Digit:=jsonReq["digit"]//客户端返回的Digit
	fmt.Println(captchaID,Digit)//输出接收到的信息
	//连接redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	digit, err := redis.String(c.Do("GET", captchaID))//这里从redis调用正确的digit
	if err != nil {
		fmt.Println("redis get failed:", err)
		Bytes:= "1"
		bytes:=[]byte(Bytes)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}else if digit!=Digit{
		Bytes:= "2"
		bytes:=[]byte(Bytes)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}else{
		Bytes:= "0"
		bytes:=[]byte(Bytes)
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func main() {
    http.HandleFunc("/captcha", showFormHandler)       //设置访问的路由
    http.HandleFunc("/captcha/verify",imageVerify)
	err := http.ListenAndServe("0.0.0.0:8667", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


