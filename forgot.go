package main

import (
	"html/template"
	"net/http"
)

// Description      :       handle registration
// returns          :       None
func handleForgot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, _ := template.ParseFiles("views/forgot.htm")
		f.Execute(w, nil)

	} else if r.Method == "POST" {
		r.ParseForm()
		usr := user{
			"",
			"",
			r.FormValue("emailadd"),
			"",
			"",
			"",
			0,
			0,
			"",
		}
		// Do something here
		// Sorry no more time
		processReset(usr)
		// sendEmailLink
	}
}
