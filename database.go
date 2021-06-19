package main

import (
	// "database/sql"
	"log"

	// "os"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)


// TODO : Use Structs : Change the functions accordingly.
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
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)

	if(err!=nil){
		log.Println("Error while hashing the password")
		return err
	}
	_, err = addStatement.Exec(rollno, name, 0, string(bytes))

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


func getUserInfo(rollno int,password string)(string,int,error){
	var name string
	var coin int
	var err error
	var userPass string

	getUserStatement,err := db.Prepare("SELECT * FROM students WHERE rollno=?")
	if err != nil {
		log.Println("Error preparing db Statement")
		return name,coin,err
	}
	defer getUserStatement.Close()

	err = getUserStatement.QueryRow(rollno).Scan(&rollno,&name,&coin, &userPass)

	// If no such row exists(Both rollno and password should match) Scan will throw an error.
	if(err != nil){
		log.Println("Error while getting user Information.")
		return name,coin,err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPass),[]byte(password))

	if(err != nil){
		return name,coin,err
	}
	//else return proper values
	return name,coin,err
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