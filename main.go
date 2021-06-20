// Write a program that connects to a database.
//  The database that we will be working with is SQLite.
//  On executing the program it should create a new table (say, User with two fields rollno and name).
//  Create a function that takes in new user details as arguments and adds it to the database.
//   Your program should not take any input from command line.
//  Aim to write clean and structured code.
package main

import (
	"database/sql"
	"log"
	"os"
	// "encoding/json"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// ? Standard Practice
// Global access to database
var db *sql.DB
var err error

// AIM : Clean Complete Code, Make is structured
//: step 1: Use Structs
//: step 2: Return Standard Errors : Remove Verbrose description (??)
//: step 3: Use JSON

type User struct {
	Rollno   int    `json:"rollno"`
	Name     string `json:"name"`
	Coin     int    `json:"coin"`
	Password string `json:"Password"`
}

func main() {
	//For testing purpose only
	os.Remove("students.db")
	//Create Database
	db, err = sql.Open("sqlite3", "./students.db")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Learn why you need this
	defer db.Close()
	//Create table
	createTable()
	// Update database
	// addUser( 190017, "Abhay Mishra","abcd")
	// addUser(190195,"Ashok kumar Saini","efgh")
	// addUser(190017,"Abhay","hello")
	// addUser(190064,"Gaurav Sharma","nice")
	// addUser(190034,"Hello world","cool")

	// Start Server
	startServer()

	// Display Users
	displayStudents()
}
