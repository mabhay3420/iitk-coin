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
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//For testing purpose only
	os.Remove("students.db")
	//Create Database
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil{
		log.Fatal(err)
	}
	// TODO: Learn why you need this
	defer db.Close()
	//Create table
	createTable(db)
	// Update database
	addUser(db, 190017, "Abhay Mishra")
	addUser(db,190195,"Ashok kumar Saini")
	addUser(db,190017,"Abhay")
	addUser(db,190064,"Gaurav Sharma")
	addUser(db,190034,"Hello world")

	// Display Users
	displayStudents(db)
}
