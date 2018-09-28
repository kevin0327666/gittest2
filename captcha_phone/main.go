package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "unsafe"
)
 
type JsonPostSample struct {
 
}

func messagesend(phonenumber string,body string)(string, string){
    params := make(map[string]interface{})

    //请登录zz.253.com获取API账号、密码以及短信发送的URL
    params["account"] = "N2632062"  //创蓝API账号
    params["password"] = "qMGx1pUoVO09ef" //创蓝API密码
    params["phone"] = phonenumber //手机号码

    //设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
    params["msg"] =url.QueryEscape(body)
    params["report"] = "true"
    bytesData, err := json.Marshal(params)
    if err != nil {
        fmt.Println(err.Error() )
        return "901","创蓝API错误"
    }
    reader := bytes.NewReader(bytesData)
    url := "http://smssh1.253.com/msg/send/json"  //短信发送URL
    request, err := http.NewRequest("POST", url, reader)
    if err != nil {
        fmt.Println(err.Error())
        return "900","URL错误"
    }
    request.Header.Set("Content-Type", "application/json;charset=UTF-8")
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return "902","发送头错误"
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
        return "903","未能读取到返回信息"
    }
    str := (*string)(unsafe.Pointer(&respBytes))
    fmt.Println(*str)
    var respbytes map[string]string
    json.Unmarshal(respBytes, &respbytes)
    code:=respbytes["code"]
    errorMsg:=respbytes["errorMsg"]
        return code,errorMsg
}

func  messageBasicsend(w http.ResponseWriter,r *http.Request) {
    if r.URL.Path != "/message/basicsend" {
        http.NotFound(w, r)
        return
    }
    var jsonReq map[string]string
    jsonBytes, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(jsonBytes, &jsonReq)
    phonenumber:=jsonReq["PhoneNumber"]
    body:=jsonReq["Body"]

    var jsonRep map[string]string
    jsonRep = make(map[string]string)
    jsonRep["code"],jsonRep["errorMsg"] = messagesend(phonenumber,body)
    bytes,_:= json.Marshal(jsonRep)
    w.WriteHeader(http.StatusOK)
    w.Write(bytes)
}

func main() {
    http.HandleFunc("/message/basicsend", messageBasicsend)       //设置访问的路由
    err := http.ListenAndServe("0.0.0.0:8666", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
