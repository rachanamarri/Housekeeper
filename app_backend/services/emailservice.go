package services

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"path/filepath"
)

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, data interface{}, template string) (bool, error) {
	var host_email = os.Getenv("HOST_EMAIL")
	var host_password = os.Getenv("HOST_PASSWORD")
	emailHost := "smtp.gmail.com"
	emailFrom := host_email
	emailPassword := host_password
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
