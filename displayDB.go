package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

// TODO: Replace These Three functions with One Comman function
// TODO: Which takes table name or some indenfier as input
func displayStudents() error {

	displayStatement, err := db.Prepare("SELECT * FROM students ORDER BY Name")
	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	defer displayStatement.Close()

	row, err := displayStatement.Query()
	if err != nil {
		log.Println("error Displaying Students")
		return err
	}
	defer row.Close()

	for row.Next() {
		var user User
		row.Scan(&user.Rollno, &user.Name, &user.Coin, &user.Password, &user.Role, &user.Activity)
		fmt.Println(user)
	}

	// Sanity Check
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
		// TODO: Better format timeStamp
		fmt.Println(stamp, award)
	}
	
	// Sanity Check
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

		// TODO: Format in a better way
		fmt.Println(stamp.Round(0), transfer)
	}

	// Sanity Check
	if err = row.Err(); err != nil {
		log.Println("error While reading rows")
		return err
	}

	return nil
}
