package main

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

// Description	: Sends out password reset link to email
// Returns			: None
func sendEmailReceipt(email string, typeOfEmail int) {
	//auth := smtp.PlainAuth("", EmailUserName, EmailServerPassword, EmailHostName)
	addr := "127.0.0.1:25"

	fromName := "FYLD"
	fromEmail := "mywebpage@mywebpage.com"
	toNames := []string{"Client"}
	toEmails := []string{email}
	subject := "Reset Password"
	body := "Please follow the link to reset password"

	toAddresses := []string{}
	for i := range toEmails {
		mAdd := mail.Address{}
		mAdd.Name = toNames[i]
		mAdd.Address = toEmails[i]
		toAddresses = append(toAddresses, mAdd.String())
	}

	toHeader := strings.Join(toAddresses, ", ")
	fAdd := mail.Address{}
	fAdd.Name = fromName
	fAdd.Address = fromEmail
	fromHeader := fAdd.String()
	subjectHeader := subject
	header := make(map[string]string)

	header["To"] = toHeader
	header["From"] = fromHeader
	header["Subject"] = subjectHeader
	header["Content-Type"] = `text/html; charset="UTF-8"`
	msg := ""
	for k, v := range header {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body
	bMsg := []byte(msg)

	// Send using local postfix service
	c, err := smtp.Dial(addr)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	if err = c.Mail(fromHeader); err != nil {
		panic(err)
	}
	for _, addr := range toEmails {
		if err = c.Rcpt(addr); err != nil {
			panic(err)
		}
	}
	w, err := c.Data()
	if err != nil {
		panic(err)
	}
	_, err = w.Write(bMsg)
	if err != nil {
		panic(err)
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}
	err = c.Quit()
	if err != nil {
		panic(err)
	}

}
