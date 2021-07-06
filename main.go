package main

import (
	"database/sql"
	"fmt"
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
	err = createTable()
	if err != nil {
		log.Fatal(err)
	}

	// display current information
	fmt.Println("===-------User List-------===")
	displayStudents()
	fmt.Println("===-------Award History------===")
	displayAward()
	fmt.Println("===-------Transfer History------===")
	displayTransfer()
	fmt.Println("===----------Here you go--------===")

	// Start Server
	startServer()

}
