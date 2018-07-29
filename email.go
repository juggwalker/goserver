package main

import (
	"github.com/go-gomail/gomail"
	"os"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func (this *Email) Send() {
	Log.Out = os.Stdout
	m := gomail.NewMessage()

	m.SetAddressHeader("From", "2823335096@qq.com", "service") // 发件人
	m.SetHeader("To",
		m.FormatAddress(this.To, "customer"))
	//m.SetHeader("Cc",
	//	m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
	//m.SetHeader("Bcc",
	//	m.FormatAddress("xxxx@gmail.com", "收件人")) // 暗送

	m.SetHeader("Subject", this.Subject)
	m.SetBody("text/html", this.Body)
	d := gomail.NewPlainDialer("smtp.qq.com", 465, "2823335096@qq.com", "sdxpdvktvvltdgch") // 发送邮件服务器、端口、发件人账号、发件人密码(QQ邮箱需要填授权码)
	if err := d.DialAndSend(m); err != nil {
		Log.Warn("发送失败")
		return
	}

	Log.Info("send email success")
}
