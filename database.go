package main

import (
	// "database/sql"
	"log"

	// "os"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//Create Table
func createTable() error {
	log.SetFlags(0)

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( rollno INTEGER PRIMARY KEY NOT NULL,name TEXT NOT NULL,coins INTEGER NOT NULL,password TEXT NOT NULL)")

	if err != nil {
		return err
	}

	log.Println("Create Student table....")
	createStatement.Exec()
	log.Println("Student table Created Succesfully.")

	return nil
}

// Add New Users
func addUser(rollno int, name string, password string) error {

	// Add New User
	addStatement, err := db.Prepare("INSERT INTO students ( rollno , name , coins,password) VALUES(?,?,?,?)")
	if err != nil {
		log.Println("Error preparing Statement")
		return err
	}
	log.Println("Add New User....")
	_, err = addStatement.Exec(rollno, name, 0, password)

	// Unique Constrain on rollno
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

// Display Student
func displayStudents() error {

	displayStatement, err := db.Prepare("SELECT * FROM students ORDER BY name")

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
		var rollno int
		var name string
		var coins int
		var password string

		row.Scan(&rollno, &name, &coins, &password)

		log.Println("RollNo:", rollno, "Name:", name, "Coins:", coins, "Password:", password)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("Error While reading rows")
		return err
	}

	return nil
}
