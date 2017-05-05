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

		msgMap := r.URL.Query()

		if len(msgMap) > 0 {
			if val, ok := msgMap["msg"]; ok {
				res := map[string]string{"message": val[0]}
				f.Execute(w, res)
			} else {
				f.Execute(w, nil)
			}
		} else {
			f.Execute(w, nil)
		}

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
		http.Redirect(w, r, "/forgot?msg="+usr.email, http.StatusSeeOther)
		// Do something here
		if id, code := processReset(usr); len(code) > 0 { // sendEmailLink
			//Send email link
			sendPasswordResetLink(usr.email, id, code)
			//Redirect to confirmation of email sent
			http.Redirect(w, r, "/linksent?msg=Password reset link sent to your email.", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/forgot?msg=Email does not exist.", http.StatusSeeOther)
		}

	}
}
