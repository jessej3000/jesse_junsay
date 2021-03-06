package main

import "net/http"

func main() {
	serveHTTP()
}

// Server and handlers
func serveHTTP() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/profile", handleProfile)
	http.HandleFunc("/forgot", handleForgot)
	http.HandleFunc("/googlelogin", handleGoogleLogin)
	http.HandleFunc("/resetpassword", handleResetPassword)

	http.ListenAndServe(WebHost+":"+WebServerPort, nil)
}
