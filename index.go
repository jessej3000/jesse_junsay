package main

import (
	"html/template"
	"net/http"
)

// Description    :     handles request to home page.
// returns        :     None
func handleHome(w http.ResponseWriter, r *http.Request) {
	f, _ := template.ParseFiles("views/index.htm")
	f.Execute(w, nil)
}
