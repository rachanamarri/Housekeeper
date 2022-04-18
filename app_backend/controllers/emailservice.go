package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
)

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	emailHost := "smtp.gmail.com"
	emailFrom := "yourEmail@gmail.com"
	emailPassword := "yourEmailPassword"
	emailPort := 587

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(template, data)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Test Email" + "!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%d", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("./%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
