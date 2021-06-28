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
