package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

// Description		:			registers new user into the database
// returns				:			(bool) true if successful and false if error
func registerGoogleUser(googleID string) int64 {

	db, err := sql.Open("mysql", DBUser+":"+DBPassword+"@/apidb?charset=utf8")
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO user("
	query = query + "googleacc) "
	query = query + "VALUES (?)"

	res, err := db.Exec(query, googleID)

	if err != nil {
		panic(err)
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return 0
		} else {
			return id
		}
	}

}

//  Definition			: 		Verify if google account exist and returns the id
//	returns					:			(int) if id exist returns the id and if not create the record and return the new id
func verifyIfGoogleAccountExist(googleID string) int64 {
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
		return row.Int64(0)
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
		// Fill up map values
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
func processReset(person user) (int64, string) {
	//Check if email exist
	if id := checkIfEmailExist(person.email); id > 0 { //sendEmailLink
		// Generate guid and save id
		GID := GUID()
		setCode(id, GID)
		return id, GID
	} else { //Else email not found
		return 0, ""
	}
}

// Description 		:			Check if email exist
// returns				:			(string) returns the email if exist
func checkIfEmailExist(email string) int64 {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT id FROM user WHERE email='" + email + "'")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return row.Int64(0)
	} else {
		return 0
	}
}

// Description		:			Function to create a guid
// returns				:			(string)
func GUID() string {
	UID := make([]byte, 16)
	res, err := io.ReadFull(rand.Reader, UID)

	if err != nil {
		panic(err)
	}

	if res != len(UID) {
		return ""
	}

	UID[8] = UID[8]&^0xc0 | 0x80
	UID[6] = UID[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", UID[0:4], UID[4:6], UID[6:8], UID[8:10], UID[10:])
}

// Description			:			sets the reset code
// returns					:			none
func setCode(id int64, code string) {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	query := "UPDATE user "
	query = query + "SET resetcode = '" + code + "' "
	query = query + "WHERE id = " + strconv.Itoa(int(id))

	_, res, err := db.Query(query)
	if res == nil {
		//Do nothing
	}
	if err != nil {
		panic(err)
	}
}

// Description 		:			Check if reset code exist
// returns				:			(bool) returns true or false otherwise
func checkIfCodeExist(code string) bool {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	res, err := db.Start("SELECT id FROM user WHERE resetcode='" + code + "'")
	if err != nil {
		panic(err)
	}

	row, err := res.GetRow()
	if err != nil {
		panic(err)
	}

	if len(row) > 0 {
		// Return true as success with the id number
		return true
	} else {
		return false
	}
}

// Description		:			update user password in the database
// returns				:			(bool) true if successful and false if error
func updatePassword(id string, password string) bool {
	db := mysql.New("tcp", "", DBHost+":"+DBPort, DBUser, DBPassword, DBName)

	err := db.Connect()
	if err != nil {
		//panic(err)
		return false
	}

	query := "UPDATE user SET "
	query = query + "pwd = SHA1('" + password + "'), "
	query = query + "resetcode = '' "
	query = query + "WHERE id = " + id

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
