package main

import (
	"errors"
	"log"
	"golang.org/x/crypto/bcrypt"
)

//Create Table
func createTable() error {

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( Rollno INTEGER PRIMARY KEY NOT NULL,Name TEXT NOT NULL,Coins INTEGER NOT NULL,Password TEXT NOT NULL, Role string TEXT NOT NULL, Activity INTEGER NOT NULL)")
	if err != nil {
		return err
	}

	awardStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS awards ( Time TIMESTAMP, AwardeeRollno INTEGER NOT NULL,Amount INTEGER NOT NULL)")
	if err != nil {
		return err
	}

	transferStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS transfers ( Time TIMESTAMP, SenderRollno INTEGER NOT NULL,RecieverRollno INTEGER NOT NULL,Amount INTEGER NOT NULL)")
	if err != nil {
		return err
	}

	// Create Tables
	createStatement.Exec()
	awardStatement.Exec()
	transferStatement.Exec()

	log.Println("tables ready.")
	return nil
}

// Add New Users
func addUser(user *User) error {

	// Add New User
	addStatement, err := db.Prepare("INSERT INTO students ( Rollno , Name , Coins,Password, Role, Activity) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Println("error preparing Statement")
		return err
	}

	if user.Password == "" || user.Name == "" {
		err = errors.New("empty Name/Password not allowed")
		log.Println(err)
		return err
	}

	// TODO: Storing a complete string in place of a bit
	// TODO: is weird. You can do better than this.

	user.Coin = 0 // starting point
	user.Role = "STUDENT"
	user.Activity = 0 // staring point

	// Find role : assign the most powerful role
	// student --> council core ---> admin
	for _, rollno := range COUNCIL_CORE {
		if rollno == user.Rollno {
			user.Role = "COUNCIL_CORE"
			break
		}
	}

	for _, rollno := range ADMIN {
		if rollno == user.Rollno {
			user.Role = "ADMIN"
			break
		}
	}

	// Valid input
	log.Println("add New User....")
	// ? Why 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println("error while hashing the Password")
		return err
	}

	_, err = addStatement.Exec(user.Rollno, user.Name, user.Coin, string(bytes),user.Role,user.Activity)
	// Unique Constraint on Rollno
	if err != nil {
		log.Println("unable to Add user")
		return err
	} else {
		log.Println("succesfully Added New User.")
	}

	return nil
}
func validateLogin(user *User, hasCookie bool) error {
	var err error
	var userPass string

	getUserStatement, err := db.Prepare("SELECT * FROM students WHERE Rollno=?")
	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	defer getUserStatement.Close()

	err = getUserStatement.QueryRow(user.Rollno).Scan(&user.Rollno, &user.Name, &user.Coin, &userPass, &user.Role, &user.Activity)
	// If no such row exists(Both Rollno and Password should match) Scan will throw an error.
	if err != nil {
		log.Println("error while getting user Information.")
		return err
	}

	// Validate login if cookie is available
	// ! Do not access password field though.
	if hasCookie {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPass), []byte(user.Password))
	if err != nil {
		return err
	}
	
	// Valid User
	return nil
}
