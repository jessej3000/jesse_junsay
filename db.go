package main

import (
	"fmt"
	"strconv"

	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

//  Definition			: 		Process google log in
//	returns					:			(bool) true if successful and false otherwise
func processGoogleLogIn(googleID string) int {
	return verifyIfGoogleAccountExist(googleID)
}

// Description		:			registers new user into the database
// returns				:			(bool) true if successful and false if error
func registerGoogleUser(googleID string) int {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO user("
	query = query + "username,"
	query = query + "pwd,"
	query = query + "email,"
	query = query + "fullname,"
	query = query + "address,"
	query = query + "telephone,"
	query = query + "longitude,"
	query = query + "latitude,"
	query = query + "googleacc) "
	query = query + "VALUES ('','','','','',0,0,'"
	query = query + googleID + "')"

	_, res, err := db.Query(query)
	if res == nil {
		//Do nothing
	}

	query = "SELECT LAST_INSERT_ID();"
	_, ress, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	row, err := ress.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		return row.Int(0)
	} else {
		return 0
	}
}

//  Definition			: 		Verify if google account exist and returns the id
//	returns					:			(int) if id exist returns the id and if not create the record and return the new id
func verifyIfGoogleAccountExist(googleID string) int {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT * FROM user WHERE googleacc='" + googleID + "'")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 { // If user exist return the id
		// Return true as success with the id number
		return row.Int(0)
	} else { //If user does not exist, register then return the new id
		return registerGoogleUser(googleID)
	}
}

//  Definition      :     verifyUser function to verify if user exist and password is correct in the database
//  Returns         :     (bool),(int) returns true or false, and returns id of user
func verifyUser(username string, pwd string) (bool, int) {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT * FROM user WHERE username='" + username + "' AND pwd=SHA1('" + pwd + "')")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true, row.Int(0)
	}

	// Return false if does not exist and 0
	return false, 0
}

//  Definition      :     checkUser function to verify if user exist in the database
//  Returns         :     (bool),(int) returns true or false, and returns id of user
func checkUser(username string, email string) (bool, int) {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT * FROM user WHERE username='" + username + "' OR email ='" + email + "'")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true, row.Int(0)
	}

	// Return false if user does not exist and 0
	return false, 0
}

// Description		:			registers new user into the database
// returns				:			(bool) true if successful and false if error
func registerUser(person user) bool {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	if ok, _ := checkUser(person.username, person.email); ok {
		return false
	}

	query := "INSERT INTO user("
	query = query + "username,"
	query = query + "pwd,"
	query = query + "email,"
	query = query + "fullname,"
	query = query + "address,"
	query = query + "telephone,"
	query = query + "longitude,"
	query = query + "latitude,"
	query = query + "googleacc) "
	query = query + "VALUES ('"
	query = query + person.username + "',SHA1('"
	query = query + person.password + "'),'"
	query = query + person.email + "','"
	query = query + person.fullname + "','"
	query = query + person.address + "','"
	query = query + person.telephone + "',"
	query = query + "0,"
	query = query + "0,'"
	query = query + person.googleacc + "')"

	_, res, err := db.Query(query)
	if res == nil {
		//Do nothing
	}
	if err != nil {
		panic(err)
	}

	return true
}

// Description		:			update user info in the database
// returns				:			(bool) true if successful and false if error
func updateUser(person user) bool {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		//panic(err)
		return false
	}

	query := "UPDATE user "
	query = query + "SET username = '" + person.username + "',"
	query = query + "pwd = SHA1('" + person.password + "'),"
	query = query + "email = '" + person.email + "',"
	query = query + "fullname = '" + person.fullname + "',"
	query = query + "address = '" + person.address + "',"
	query = query + "telephone = '" + person.telephone + "',"
	query = query + "longitude = 0,"
	query = query + "latitude = 0,"
	query = query + "googleacc = '' "
	query = query + "WHERE id = " + strconv.Itoa(MyID)

	fmt.Println(query)
	_, res, err := db.Query(query)
	if res == nil {
		//Do nothing
	}
	if err != nil {
		//panic(err)
		return false
	}

	return true
}

// Description		:			gets user info from db
// returns				:			(user struct)
func getUser(id int) map[string]string {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	query := "SELECT "
	query = query + "username, "
	query = query + "SHA1(pwd) AS pwd, "
	query = query + "email, "
	query = query + "fullname, "
	query = query + "address, "
	query = query + "telephone, "
	query = query + "longitude, "
	query = query + "latitude, "
	query = query + "googleacc "
	query = query + "FROM user WHERE id=" + strconv.Itoa(id)

	res, err := db.Start(query)
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	person := make(map[string]string)

	if len(row) > 0 {
		// Return true as success with the id number
		person["username"] = row.Str(0)
		person["password"] = row.Str(1)
		person["email"] = row.Str(2)
		person["fullname"] = row.Str(3)
		person["address"] = row.Str(4)
		person["telephone"] = row.Str(5)
		person["longitude"] = row.Str(6)
		person["latitude"] = row.Str(7)
		person["googleacc"] = row.Str(8)
	} else {
		return nil
	}

	return person
}

// Description		:			process reset link for password
// returns				:			none
func processReset(person user) {
	return
}
