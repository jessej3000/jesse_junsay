package main

import "net/http"

func main() {
	serveHTTP()
}

// Server
func serveHTTP() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/profile", handleProfile)
	http.HandleFunc("/forgot", handleForgot)

	http.ListenAndServe("localhost:"+WebServerPort, nil)
}
