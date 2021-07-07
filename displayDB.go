package main

import (
	// "database/sql"
	// "errors"
	"fmt"
	"log"
	"time"

	// "os"
	// "context"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)

// Display Student
func displayStudents() error {

	displayStatement, err := db.Prepare("SELECT * FROM students ORDER BY Name")

	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	// TODO: Learn More
	defer displayStatement.Close()
	row, err := displayStatement.Query()

	if err != nil {
		log.Println("error Displaying Students")
		return err
	}
	// TODO: Learn More
	defer row.Close()
	for row.Next() {
		var user User

		row.Scan(&user.Rollno, &user.Name, &user.Coin, &user.Password, &user.Role, &user.Activity)

		fmt.Println(user)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("error While reading rows")
		return err
	}

	return nil
}

func displayAward() error {
	displayStatement, err := db.Prepare("SELECT * FROM awards ORDER BY Time ")

	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	// TODO: Learn More
	defer displayStatement.Close()
	row, err := displayStatement.Query()

	if err != nil {
		log.Println("error Displaying Awards")
		return err
	}
	defer row.Close()
	for row.Next() {
		var award awardRequest
		var stamp time.Time
		row.Scan(&stamp, &award.Rollno, &award.Award)

		// we do not need time here.
		fmt.Println(stamp, award)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("error While reading rows")
		return err
	}

	return nil

}
func displayTransfer() error {
	displayStatement, err := db.Prepare("SELECT * FROM transfers ORDER BY Time ")

	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	// TODO: Learn More
	defer displayStatement.Close()
	row, err := displayStatement.Query()

	if err != nil {
		log.Println("error Displaying Transfers")
		return err
	}
	defer row.Close()
	for row.Next() {
		var transfer transferRequest
		var stamp time.Time
		row.Scan(&stamp, &transfer.FromRollno, &transfer.ToRollno, &transfer.Amount)

		// no need of log here
		fmt.Println(stamp.Round(0), transfer)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("error While reading rows")
		return err
	}

	return nil

}
