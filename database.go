package main

import (
	// "database/sql"
	"errors"
	"log"
	// "os"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// TODO : Use Structs : Change the functions accordingly.
//Create Table
func createTable() error {

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( Rollno INTEGER PRIMARY KEY NOT NULL,Name TEXT NOT NULL,Coins INTEGER NOT NULL,Password TEXT NOT NULL)")

	if err != nil {
		return err
	}

	// log.Println("Create Student table....")
	createStatement.Exec()
	log.Println("Tables ready.")

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
	log.Println("Add New User....")
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("error while hashing the Password")
		return err
	}
	_, err = addStatement.Exec(user.Rollno, user.Name, user.Coin, string(bytes))

	// Unique Constrain on Rollno
	if err != nil {
		log.Println("Unable to Add user")
		return err
	} else {
		log.Println("Succesfully Added New User.")
	}

	// testing purpose
	displayStudents()

	return nil
}

func getUserInfo(user *User) error {
	var err error
	var userPass string

	getUserStatement, err := db.Prepare("SELECT * FROM students WHERE Rollno=?")
	if err != nil {
		log.Println("Error preparing db Statement")
		return err
	}
	defer getUserStatement.Close()

	err = getUserStatement.QueryRow(user.Rollno).Scan(&user.Rollno, &user.Name, &user.Coin, &userPass)

	// If no such row exists(Both Rollno and Password should match) Scan will throw an error.
	if err != nil {
		log.Println("Error while getting user Information.")
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPass), []byte(user.Password))

	if err != nil {
		return err
	}
	//else return proper values
	return err
}

// Display Student
func displayStudents() error {

	displayStatement, err := db.Prepare("SELECT * FROM students ORDER BY Name")

	if err != nil {
		log.Println("Error preparing db Statement")
		return err
	}
	// TODO: Learn More
	defer displayStatement.Close()
	row, err := displayStatement.Query()

	if err != nil {
		log.Println("Error Displaying Students")
		return err
	}
	// TODO: Learn More
	defer row.Close()
	for row.Next() {
		var user User

		row.Scan(&user.Rollno, &user.Name, &user.Coin, &user.Password)

		log.Println("User:", "Rollno:", user.Rollno, "Name:", user.Name)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("Error While reading rows")
		return err
	}

	return nil
}
