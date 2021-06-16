package main

import (
	"database/sql"
	"log"

	// "os"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
)
//Create Table
func createTable(db *sql.DB) {
	log.SetFlags(0)

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( rollno INT PRIMARY KEY,name TEXT)")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Create Student table....")
	createStatement.Exec()
	log.Println("Student table Created Succesfully.")
}

// Add New Users
func addUser(db *sql.DB, rollno int, name string) {

	// Add New User
	addStatement, err := db.Prepare("INSERT INTO students ( rollno , name ) VALUES(?,?)")
	if err != nil {
		log.Fatalln("adding failed:", err)
	}
	log.Println("Add New User....")
	_, err = addStatement.Exec(rollno, name)

	// Unique Constrain on rollno
	if err != nil {
		log.Println("Cannot Add User:", err)
	} else {
		log.Println("Succesfully Added New User.")
	}
}

// Display Student
func displayStudents(db *sql.DB) {

	displayStatement, err := db.Prepare("SELECT * FROM students ORDER BY name")

	if err != nil {
		log.Fatalln("displayStudents:", err)
	}
	// TODO: Learn More
	defer displayStatement.Close()
	row, err := displayStatement.Query()

	if err != nil {
		log.Fatalln("displayStudents:", err)
	}
	// TODO: Learn More
	defer row.Close()
	for row.Next() {
		var rollno int
		var name string

		row.Scan(&rollno, &name)

		log.Println("RollNo:", rollno, "Name:", name)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Fatalln("displayStudents:", err)
	}
}
