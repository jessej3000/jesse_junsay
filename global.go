package main

const (
	// DBUser MySQL database username
	DBUser = "root" //"test" //
	// DBPassword MySQL database password
	DBPassword = "" //"root" //"test" //
	// DBName database name
	DBName = "apidb"
	// DBHost MySQL host server
	DBHost = "localhost"
	// DBPort MySQL port
	DBPort = "3306" //"8889"
	// WebServerPort for golang web server
	WebServerPort = "8080"
	// WebHost for golang web server
	WebHost = "172.31.42.235"
)

type result struct {
	message string
}

type user struct {
	username  string
	password  string
	email     string
	fullname  string
	address   string
	telephone string
	longitude float64
	latitude  float64
	googleacc string
}

// MyID global id of the user will later be set during login
var MyID = 0
var ResetID = ""
