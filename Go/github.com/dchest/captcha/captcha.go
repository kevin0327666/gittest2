package captcha

import (
	"bytes"
	"io"
	"encoding/base64"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 4
)

// Creator: chending
// NewId creates a new captcha with the standard length, saves it in the internal
// storage and returns its id and its digit.
func New() (string, string) {
	return NewLen(DefaultLen)
}


func NewLen(length int) (result, captchaId string) {
	digitBytes := RandomDigits(length)
	captchaid := RandomDigits(16)
	var content bytes.Buffer
	WriteImage(&content, digitBytes, StdWidth, StdHeight)
	result ="data:image/png;base64,"
	result += base64.StdEncoding.EncodeToString(content.Bytes())
	var digit string
	for _, d := range digitBytes {
		digit += string('0' + d)//被发送
	}
	for _, d := range captchaid {
		captchaId += string('0' + d)//被发送
	}

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", captchaId, digit,"EX", "1200")//发送的验证码60秒会失效
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	return
}



// WriteImage writes PNG-encoded image representation of the captcha with the
// given id. The image will have the given width and height.
func WriteImage(w io.Writer, digit []byte, width, height int) error {
	_, err := NewImage(digit, width, height).WriteTo(w)
	return err
}

