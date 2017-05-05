package main

import (
	"html/template"
	"net/http"
)

// Description      :       handle registration
// returns          :       None
func handleProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("logged")
	if err == nil { // Check if cookie value is logged = 1
		if cookie.Value == "1" {
			if r.Method == "GET" {
				f, _ := template.ParseFiles("views/profile.htm")

				person := getUser(MyID)
				f.Execute(w, person)

			} else if r.Method == "POST" {
				r.ParseForm()
				usr := user{
					r.FormValue("username"),
					r.FormValue("password"),
					r.FormValue("emailadd"),
					r.FormValue("fullname"),
					r.FormValue("address"),
					r.FormValue("telephone"),
					r.FormValue("lon"),
					r.FormValue("lat"),
					"",
				}
				//Check if username and password are correct
				if updateUser(usr) {
					http.Redirect(w, r, "/dashboard?msg=Profile updated.", http.StatusSeeOther)
				} else {
					http.Redirect(w, r, "/dashboard?msg=Sorry, something went wrong during the update. Please try again.", http.StatusSeeOther)
				}
			}
		} else { // redirect to login
			http.Redirect(w, r, "/login?msg=You need to login", http.StatusSeeOther)
		}
	} else { // redirect to login
		http.Redirect(w, r, "/login?msg=You need to login", http.StatusSeeOther)
	}
}
