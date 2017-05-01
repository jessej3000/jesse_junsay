package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Description      :     Handles request for landing page dashboard
// Returns          :     None
func handleDashboard(w http.ResponseWriter, r *http.Request) {
	//Check cookie
	cookie, err := r.Cookie("logged")
	fmt.Println("Got cookie")
	if err == nil { // Check if cookie value is logged = 1
		fmt.Println("Got cookie if not nil:" + cookie.Value)
		if cookie.Value == "1" {
			f, _ := template.ParseFiles("views/dashboard.htm")

			f.Execute(w, nil)
		} else { // redirect to login
			http.Redirect(w, r, "/login?msg=You need to login", http.StatusSeeOther)
		}
	} else { // redirect to login
		http.Redirect(w, r, "/login?msg=You need to login", http.StatusSeeOther)
	}
}
