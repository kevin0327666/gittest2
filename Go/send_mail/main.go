package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func mailBasicsend(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/mail/basicsend" {
		http.NotFound(w, r)
		return
	}
	var jsonReq map[string]string
	jsonBytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(jsonBytes, &jsonReq)
	sento:=jsonReq["sendto"]
	nickname:=jsonReq["nickname"]
	subject:=jsonReq["subject"]
	body:=jsonReq["body"]

	auth := smtp.PlainAuth("", "kevin0327666@163.com", "kevin0327666", "smtp.163.com")
	user := "kevin0327666@163.com"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	to := []string{sento}//发送对象
//	nickname = "test"//用户名
//	subject = "test mail"//邮件主题//这里有问题
//	body = "This is the email body."//邮件内容
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.163.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		bytes,_:= json.Marshal("1")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}else {
		bytes,_:= json.Marshal("0")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}
func main() {
	http.HandleFunc("/mail/basicsend", mailBasicsend)       //设置访问的路由
	err := http.ListenAndServe("192.168.40.49:9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}