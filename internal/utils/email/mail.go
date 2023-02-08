package email

import (
	"selefra-demo/internal/model"
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string // 发件人
	MailPass string // 发件人密码
	MailTo   string // 收件人 多个用,分割
	Subject  string // 邮件主题
	Body     string // 邮件内容
}

func Send(o *Options) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", o.MailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	//设置邮件正文
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	return d.DialAndSend(m)
}

// 返回要发送u邮件的users
func ConvertToUsers(user []model.User) string {
	var usersLink string
	for i := 0; i < len(user); i++ {
		usersLink += user[i].EmailLink + ","
	}
	usersLink = usersLink[:len(usersLink)-1]
	return usersLink
}
