package main

import (
	"html/template"
	"net/http"
)

// Description    :     handles request to home page.
// returns        :     None
func handleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("logged")
	if err == nil { // if cookie exist
		cookie.Value = "0"
		cookie.MaxAge = -1
		http.SetCookie(w, cookie) // destroy cookie by setting it to 0
	}

	f, _ := template.ParseFiles("views/index.htm")
	f.Execute(w, nil)
}
