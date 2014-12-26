package main

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

var (
	user     string = "your@beequick.cn"
	password string = "password"
	host     string = "smtp.exmail.qq.com"
	port     string = "25"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) string {
	return coder.EncodeToString(src)
}

func SendMail(to []string, title string, body string) error {

	auth := smtp.PlainAuth("", user, password, host)

	header := make(map[string]string)
	header["From"] = user
	header["To"] = strings.Join(to, ";")
	header["Subject"] = "Online Error by " + title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=UTF-8"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body = strings.Replace(strings.Replace(strings.Replace(strings.Replace(body, "\n\r", "\r", -1), "\r\n", "\n", -1), "\r", "\n", -1), "\n", "\r\n", -1)

	//body = "<html><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><title>" + title + "</title></head><body>" + body + "</body></html>"

	message += fmt.Sprintf("\r\n%s", body)

	//fmt.Print(message)
	//return nil

	err := smtp.SendMail(host+":"+port, auth, user, to, []byte(message))
	return err
}
