package main

import (
	"html/template"
	"net/http"
)

//  Description     :     Function to handle google login api
//  returns         :     none
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		if r.FormValue("loggedin") == "true" {
			if id := verifyIfGoogleAccountExist(r.FormValue("googleID")); id > 0 {
				//log in and redirect to landing page
				//Check cookie
				MyID = id
				cookie, err := r.Cookie("logged")
				if err == http.ErrNoCookie { // if not then create one
					cookie = &http.Cookie{
						Name:   "logged",
						Value:  "1",
						MaxAge: 300,
					}
				} else { // if cookie exist set value = 1
					cookie.Value = "1"
					cookie.MaxAge = 300
				}

				// set cookie
				http.SetCookie(w, cookie)

				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			} else {
				f, _ := template.ParseFiles("views/sorry.htm")
				f.Execute(w, nil)
			}
		} else {
			f, _ := template.ParseFiles("views/sorry.htm")
			f.Execute(w, nil)
		}
	} else {
		f, _ := template.ParseFiles("views/sorry.htm")
		f.Execute(w, nil)
	}
}
