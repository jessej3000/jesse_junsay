package main

import (
	"html/template"
	"net/http"
)

// Description      :       handle registration
// returns          :       None
func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, _ := template.ParseFiles("views/register.htm")
		msgMap := r.URL.Query()

		if len(msgMap) > 0 {
			if val, ok := msgMap["msg"]; ok {
				//res := result{message: val[0]}
				res := map[string]string{"message": val[0]}
				//fmt.Print(res.message)
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
		if registerUser(usr) {
			http.Redirect(w, r, "/login?msg=Congratulation! You are registered", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/register?msg=Sorry, something went wrong during registration.", http.StatusSeeOther)
		}
	}
}
