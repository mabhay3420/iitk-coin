package main

import (
	// "database/sql"
	"errors"
	// "fmt"
	"log"
	"time"

	// "os"
	"context"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

//Create Table
func createTable() error {

	// Roll Number Should Be unique.
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS students ( Rollno INTEGER PRIMARY KEY NOT NULL,Name TEXT NOT NULL,Coins INTEGER NOT NULL,Password TEXT NOT NULL)")
	if err != nil {
		return err
	}
	awardStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS awards ( Time TIMESTAMP, AwardeeRollno INTEGER NOT NULL,Amount INTEGER NOT NULL)")
	if err != nil {
		return err
	}
	transferStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS transfers ( Time TIMESTAMP, SenderRollno INTEGER NOT NULL,RecieverRollno INTEGER NOT NULL,Amount INTEGER NOT NULL)")
	if err != nil {
		return err
	}
	// Create Tables
	createStatement.Exec()
	awardStatement.Exec()
	transferStatement.Exec()

	log.Println("tables ready.")
	return nil
}

// Add New Users
func addUser(user *User) error {

	// Add New User
	addStatement, err := db.Prepare("INSERT INTO students ( Rollno , Name , Coins,Password) VALUES(?,?,?,?)")

	if err != nil {
		log.Println("error preparing Statement")
		return err
	}

	if user.Password == "" || user.Name == "" {
		err = errors.New("empty Name/Password not allowed")
		log.Println(err)
		return err
	}

	// Valid input
	log.Println("add New User....")
	// ? Why 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("error while hashing the Password")
		return err
	}
	_, err = addStatement.Exec(user.Rollno, user.Name, user.Coin, string(bytes))

	// Unique Constraint on Rollno
	if err != nil {
		log.Println("unable to Add user")
		return err
	} else {
		log.Println("succesfully Added New User.")
	}

	// testing purpose
	// displayStudents()

	return nil
}
func validateLogin(user *User, hasCookie bool) error {
	var err error
	var userPass string
	getUserStatement, err := db.Prepare("SELECT * FROM students WHERE Rollno=?")
	if err != nil {
		log.Println("error preparing db Statement")
		return err
	}
	defer getUserStatement.Close()

	err = getUserStatement.QueryRow(user.Rollno).Scan(&user.Rollno, &user.Name, &user.Coin, &userPass)

	// If no such row exists(Both Rollno and Password should match) Scan will throw an error.
	if err != nil {
		log.Println("error while getting user Information.")
		return err
	}

	// must be true
	// ! Do not access password field though.
	if hasCookie {
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPass), []byte(user.Password))

	if err != nil {
		return err
	}
	//else return proper values
	return err
}
func awardCoin(award *awardRequest) error {

	// Need to complete this thing in one go.
	ctx := context.Background()

	// ? transactions options
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// ignore rollback if the tx has been commited later in the function
	defer tx.Rollback()

	awardStatement, err := tx.Prepare("UPDATE students SET Coins = Coins + ? WHERE Rollno = ? ")
	if err != nil {
		log.Println("error while preparing award statement")

		tx.Rollback()
		return err
	}
	defer awardStatement.Close()

	_, err = awardStatement.Exec(award.Award, award.Rollno)
	if err != nil {
		log.Println("error while awarding the student")

		tx.Rollback()
		return err
	}
	// Update award table
	recordStatement, err := tx.Prepare("INSERT INTO awards ( Time, AwardeeRollno, Amount ) VALUES(?,?,?)")
	if err != nil {
		log.Println("error while preparing award statement")

		tx.Rollback()
		return err
	}
	defer recordStatement.Close()

	_, err = recordStatement.Exec(time.Now(), award.Rollno, award.Award)
	if err != nil {
		log.Println("error while recording the award the student")

		tx.Rollback()
		return err
	}

	// Succesful

	err = tx.Commit()
	// displayAward()

	return err
}

func transferCoin(transfer *transferRequest) error {

	// Need to complete this thing in one go.
	ctx := context.Background()

	// ? transactions options
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// ignore rollback if the tx has been commited later in the function
	defer tx.Rollback()

	// Update sender balance
	senderStatement, err := tx.Prepare("UPDATE students SET Coins = Coins - ? WHERE Rollno = ? ")
	if err != nil {
		log.Println("error while preparing award statement for sender")

		tx.Rollback()
		return err
	}
	defer senderStatement.Close()

	_, err = senderStatement.Exec(transfer.Amount, transfer.FromRollno)
	if err != nil {
		log.Println("error while updating sender info")

		tx.Rollback()
		return err
	}
	// update Reciever balance
	recieverStatement, err := tx.Prepare("UPDATE students SET Coins = Coins + ? WHERE Rollno = ? ")
	if err != nil {
		log.Println("error while preparing statement for reciever")

		tx.Rollback()
		return err
	}
	defer recieverStatement.Close()

	_, err = recieverStatement.Exec(transfer.Amount, transfer.ToRollno)
	if err != nil {
		log.Println("error while updating reciever info")

		tx.Rollback()
		return err
	}
	// Update transfer table
	recordStatement, err := tx.Prepare("INSERT INTO transfers ( Time , SenderRollno ,RecieverRollno , Amount ) VALUES(?,?,?,?)")
	if err != nil {
		log.Println("error while preparing record statement of transfer")

		tx.Rollback()
		return err
	}
	defer recordStatement.Close()

	log.Println(transfer)
	_, err = recordStatement.Exec(time.Now(), transfer.FromRollno, transfer.ToRollno, transfer.Amount)

	if err != nil {
		log.Println("error while recording the transfer")

		tx.Rollback()
		return err
	}

	// Succesful

	err = tx.Commit()
	// displayAward()

	return err
}

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

		row.Scan(&user.Rollno, &user.Name, &user.Coin, &user.Password)

		log.Println("User:", "Rollno:", user.Rollno, "Name:", user.Name)
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

		log.Println("Time:", stamp, "Rollno:", award.Rollno, "Award :", award.Award)
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

		log.Println("Time:", stamp, "Sender Rollno:", transfer.FromRollno, "Reciever Rollno:", transfer.ToRollno, "Amount:", transfer.Amount)
	}
	// Maybe not in the right
	// format while implicit conversion (e.g. String to Int)
	if err = row.Err(); err != nil {
		log.Println("error While reading rows")
		return err
	}

	return nil

}
