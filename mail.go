package main
import (
	"net/smtp"
	"fmt"
	"strings"
	"encoding/base64"
)

/*
 *	user : example@example.com login smtp server user
 *	password: xxxxx login smtp server password
 *	host: smtp.example.com:port   smtp.163.com:25
 *	to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */


func SendMail(user, password, host, to, subject, body, mailtype string) error{
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/"+ mailtype + "; charset=UTF-8"
	}else{
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<"+ user +">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
//	user := "zhaoming.14254376@163.com"
	user := "zhaoming@hoolai.com"
	passwd_en := "YzcyM294"
	p , err := base64.StdEncoding.DecodeString(passwd_en)
	println(err)
	if err != nil {
		println("error: conn mail server error")
		return
	}
	password := string(p)

//	host := "smtp.163.com:25"
	host := "smtp.hoolai.com:25"
	to := "zhaoming200808@gmail.com;zhaoming@hoolai.com"

	subject := "焚天开服检查"

	body := `
	<html>
	<body>
	<h1> 焚天自动添加监控信息 </h1>
	<h3> 自动添加zabbix监控：y </h3>
	<h3> 理论剩余新服数量：  10 </h3>
	</body>
	</html>
	`
	fmt.Println("send email")
	err = SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	}else{
		fmt.Println("send mail success!")
	}
}

