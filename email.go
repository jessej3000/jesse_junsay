package main

import (
	"net/smtp"
	"strconv"
)

// Description	: Sends smtp email using google smtp server
// Returns			: None
func sendPasswordResetLink(email string, id int64, code string) {
	from := "jessejmwp2017@gmail.com"
	pass := "20mwp17golang"

	to := email
	href := "http://ec2-52-36-191-60.us-west-2.compute.amazonaws.com:8080/resetpassword?id=" + strconv.Itoa(int(id)) + "&code=" + code
	link := "<a href='" + href + "'>" + href + "</>"
	body := "<html><body>Please follow the link to reset password: " + link + "</body></html>"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: Password Reset Link\n\n" + body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		panic(err)
	}
}
