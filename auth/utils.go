package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/smtp"
	"time"

	"github.com/McaxDev/backend/misc/rpc"
)

func AuthCode(number, authcode string, data MsgSent) (bool, error) {

	msgSent, exists := data[number]
	if !exists {
		return false, fmt.Errorf("请先申请验证码：%s\n", number)
	}

	if expiry := msgSent.Expiry; time.Now().After(expiry) {
		return false, fmt.Errorf("验证码已过期%v\n", expiry)
	}

	if authcode != msgSent.Authcode {
		return false, fmt.Errorf("验证码不正确：%s\n", authcode)
	}

	return true, nil
}

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
		config.SMTP.Server+":"+config.SMTP.Port,
		smtp.PlainAuth("",
			config.SMTP.Mail,
			config.SMTP.Password,
			config.SMTP.Server,
		),
		config.SMTP.Mail,
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

	response, err := MiscClient.GetGeoByIP(
		context.Background(), &rpc.String{Data: clientIp},
	)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("/data/email.html")
	if err != nil {
		return err
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
		Location: response.Country + response.City,
	}); err != nil {
		return err
	}

	if err := SendEmail(email, "验证码邮件", buffer.Bytes()); err != nil {
		return err
	}

	return nil
}

func ClearSent(datas ...MsgSent) {
	for _, data := range datas {
		for key, value := range data {
			if time.Now().After(value.Expiry) {
				delete(data, key)
			}
		}
	}
}
