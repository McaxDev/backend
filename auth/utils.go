package main

import (
	"bytes"
	"net/smtp"
	"text/template"
	"time"

	"github.com/McaxDev/backend/utils"
)

func SendEmail(email, title string, content []byte) error {

	var buffer bytes.Buffer
	buffer.Write([]byte(
		"From: Axolotland Gaming Club <axolotland@163.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + title + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n",
	))
	buffer.Write(content)

	if err := smtp.SendMail(
		Config.SMTP.Server+":"+Config.SMTP.Port,
		smtp.PlainAuth("",
			Config.SMTP.Mail,
			Config.SMTP.Password,
			Config.SMTP.Server,
		),
		Config.SMTP.Mail,
		[]string{email},
		buffer.Bytes(),
	); err != nil {
		return err
	}
	return nil
}

func SendEmailCode(
	email, code, clientIp string, expiry time.Time,
) error {

	tmpl, err := template.ParseFiles("/data/email.html")
	if err != nil {
		return err
	}

	if Config.GeoSrvAddr != "" {
		resp, err := utils.Get[map[string]string](Config.GeoSrvAddr)
		if err != nil {
			return err
		}
		clientIp = (*resp)["province"] + (*resp)["city"]
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, &struct {
		Email    string
		Authcode string
		Expiry   string
		Location string
	}{
		Email:    email,
		Authcode: code,
		Expiry:   expiry.Format("2006-01-02 15:04:05"),
		Location: clientIp,
	}); err != nil {
		return err
	}

	if err := SendEmail(email, "验证码邮件", buffer.Bytes()); err != nil {
		return err
	}

	return nil
}
