package main

import (
	"html/template"
	"net/http"
)

// Description      :     Login handler
// returns          :     None
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, _ := template.ParseFiles("views/login.htm")
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
		//Check if username and password are correct
		result, id := verifyUser(r.FormValue("username"), r.FormValue("password"))
		MyID = id
		if result { // Login if true
			//Check cookie
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
			http.Redirect(w, r, "/login?msg=Invalid Username or Password", http.StatusSeeOther)
		}
	}
}
