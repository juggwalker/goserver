package main

import (
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
	"os"
)

func SendMail() {
	log.Out = os.Stdout
	m := gomail.NewMessage()

	m.SetAddressHeader("From", "2823335096@qq.com", "service") // 发件人

	m.SetHeader("To",
		m.FormatAddress("23335096@qq.com", "jugg")) // 收件人
	//m.SetHeader("Cc",
	//	m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
	//m.SetHeader("Bcc",
	//	m.FormatAddress("xxxx@gmail.com", "收件人")) // 暗送

	m.SetHeader("Subject", "hornet Title") // 主题

	m.SetBody("text/html", "<p>xxxxx 我是正文</p>") // 可以放html..还有其他的

	//m.Attach("我是附件") //添加附件

	d := gomail.NewPlainDialer("smtp.qq.com", 465, "2823335096@qq.com", "sdxpdvktvvltdgch") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.WithFields(logrus.Fields{
			"status": "fail",
		}).Info("发送失败")

		return
	}

	log.WithFields(logrus.Fields{
		"status": "success",
	}).Info("email is yet!")
}
