package main

import "net/smtp"

// Description	: Sends out password reset link to email
// Returns			: None
func sendPasswordResetLink(email string, id int64, code string) {
	from := "jessejmwp2017@gmail.com"
	pass := "20mwp17golang"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	to := email
	href := "http://view-source:ec2-52-36-191-60.us-west-2.compute.amazonaws.com:8080/resetpassword?id=" + string(id) + "&code=" + code
	link := "<a href='" + href + "'>" + href + "</>"
	body := "Please follow the link to reset password: " + link

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Password Reset Link\n\n" + mime + body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		panic(err)
	}
}
