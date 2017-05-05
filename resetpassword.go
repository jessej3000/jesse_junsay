package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Description      :     Login handler
// returns          :     None
func handleResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, _ := template.ParseFiles("views/resetpassword.htm")
		msgMap := r.URL.Query()

		if len(msgMap) > 0 {
			if val, ok := msgMap["code"]; ok { // Check if code exist
				// Check if code exist
				fmt.Println("CODE=====" + val[0])
				if p := checkIfCodeExist(val[0]); p { //Show reset view
					fmt.Println(p)
					ResetID = msgMap["id"][0]
					f.Execute(w, nil)
				} else {
					http.Redirect(w, r, "/login?msg=Reset password link no longer valid.", http.StatusSeeOther)
				}
			} else {
				http.Redirect(w, r, "/login?msg=Reset password link no longer valid.", http.StatusSeeOther)
			}
		} else {
			http.Redirect(w, r, "/login?msg=Reset password link no longer valid.", http.StatusSeeOther)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		if updatePassword(ResetID, r.FormValue("password")) { // Login if true
			http.Redirect(w, r, "/login?msg=Password reset successful", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login?msg=Unable to reset password", http.StatusSeeOther)
		}
	}
}
