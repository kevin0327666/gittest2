FROM ubuntu

MAINTAINER Kevin "2711249603@qq.com"

#install golang
RUN apt-get update && apt-get install -y \
		golang\
		redis-server

#config GOROOT
ENV GOROOT /usr/lib/go
ENV PATH=$PATH:/usr/lib/golang/bin

# config GOPATH
RUN mkdir -p /root/gopath
RUN mkdir -p /root/gopath/src
RUN mkdir -p /root/gopath/pkg
RUN mkdir -p /root/gopath/bin
ENV GOPATH /root/gopath

# copy source files
RUN mkdir -p /root/gopath/src/captcha-kevin
ADD . /root/gopath/src/captcha-kevin
# add the lib
ADD . /root/gopath/src


#build the captcha service
WORKDIR /root/gopath/src/captcha-kevin
RUN go build -o messagemain main.go
RUN chmod +x ./test.sh && \
    chmod +x ./messagemain

#expose
EXPOSE 8667

#start redis and captcha.
ENTRYPOINT ["./test.sh"]

