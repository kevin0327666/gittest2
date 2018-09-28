短信发送服务：

step1：
main.go
修改监听地址0.0.0.0-->本机ip

step2:
#构建docker image
在当前文件夹开终端，运行
docker build -t message .

step3:
#构建容器
docker run -p 8666:8666 -d message

END:
#接口详见接口文档。