//  Write a program that connects to a database.
//  The database that we will be working with is SQLite.
//  On executing the program it should create a new table (say, User with two fields rollno and name).
//  Create a function that takes in new user details as arguments and adds it to the database.
//   Your program should not take any input from command line.
//  Aim to write clean and structured code.
package main

import (
	"database/sql"
	"log"
	// "os"
	// "encoding/json"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

type User struct {
	Rollno   int    `json:"rollno"`
	Name     string `json:"name"`
	Coin     int    `json:"coin"`
	Password string `json:"password"`
}


//TODO : 1. Errors in json format 2. JWT 3. secret endpoint
func main() {
	//For testing purpose only
	// os.Remove("students.db")
	//Create Database
	db, err = sql.Open("sqlite3", "./students.db")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Learn why you need this
	defer db.Close()

	//Create table
	createTable()

	// Start Server
	startServer()

	// Display Users
	displayStudents()
}
