package main

import (
	// "database/sql"
	"errors"
	// "fmt"
	"log"
	// "time"

	// "os"
	// "context"
	// Import sqlite3
	// _ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

//Create Table
func createTable() error {

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( Rollno INTEGER PRIMARY KEY NOT NULL,Name TEXT NOT NULL,Coins INTEGER NOT NULL,Password TEXT NOT NULL)")
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
	addStatement, err := db.Prepare("INSERT INTO students ( Rollno , Name , Coins,Password) VALUES(?,?,?,?)")

	if err != nil {
		log.Println("error preparing Statement")
		return err
	}

	if user.Password == "" || user.Name == "" {
		err = errors.New("empty Name/Password not allowed")
		log.Println(err)
		return err
	}

	// Valid input
	log.Println("add New User....")
	// ? Why 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("error while hashing the Password")
		return err
	}
	_, err = addStatement.Exec(user.Rollno, user.Name, user.Coin, string(bytes))

	// Unique Constraint on Rollno
	if err != nil {
		log.Println("unable to Add user")
		return err
	} else {
		log.Println("succesfully Added New User.")
	}

	// testing purpose
	// displayStudents()

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

	err = getUserStatement.QueryRow(user.Rollno).Scan(&user.Rollno, &user.Name, &user.Coin, &userPass)

	// If no such row exists(Both Rollno and Password should match) Scan will throw an error.
	if err != nil {
		log.Println("error while getting user Information.")
		return err
	}

	// must be true
	// ! Do not access password field though.
	if hasCookie {
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPass), []byte(user.Password))

	if err != nil {
		return err
	}
	//else return proper values
	return nil
}