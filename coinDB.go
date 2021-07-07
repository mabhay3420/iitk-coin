package main

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"log"
	"time"

	// "os"
	"context"
	// Import sqlite3
	_ "github.com/mattn/go-sqlite3"
	// "golang.org/x/crypto/bcrypt"
)
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