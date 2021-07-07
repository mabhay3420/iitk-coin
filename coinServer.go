// Task

// 1. You have already created coin related endpoints,
//    now you have to secure them and set relevant permissions.
//    Make sure all the secure endpoints are accessible only via a valid JWT token,
//    and also the owner of the JWT token should have permission to do what they doing.
//    Modify the coin related endpoints such that they comply with the
//    guidelines given in the project idea. For example, user roles,
//    tax on transfers, rules on who cannot earn, etc. Just go through
//    the doc and make sure you have all of them.
//    You can make your own design choices to implement this.

// 2. Add new tables and maintain history of transactions (transfers and rewards).

package main

import (
	"encoding/json"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	// "log"
	"net/http"
	// "time"
	// "strconv"
)

/* GLOBAL VARIABLES*/
var MAX_COIN int = 100000 // max no of coins a person can hold
/* GLOBAL VARIABLES*/
type balanceRequest struct {
	Rollno int `json:"rollno"`
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	// Must be a get request
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		fmt.Println("Unsupported Method ")
		return
	}

	err := checkCookie(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// * Authorized
	var balance balanceRequest
	if err := json.NewDecoder(r.Body).Decode(&balance); err != nil {
		http.Error(w, "Invalid Input to the form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	user := User{Rollno: balance.Rollno}
	err = validateLogin(&user, true) // has a cookie
	if err != nil {
		// extra sanity check
		fmt.Println("Invalid id/passoword")
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}
	response := Response{Rollno: user.Rollno, Name: user.Name, Coin: user.Coin}

	fmt.Println("Balance of", user.Name, "is", user.Coin)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type awardRequest struct {
	Rollno int `json:"rollno"`
	Award  int `json:"award"`
}

func awardHandler(w http.ResponseWriter, r *http.Request) {
	// Must be a post request
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		fmt.Println("Unsupported Method Name")
		return
	}

	if err := checkCookie(w, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// * Authorized
	var award awardRequest
	if err := json.NewDecoder(r.Body).Decode(&award); err != nil {
		http.Error(w, "Invalid Input to the form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// All information provided
	// awardUser can do both the things decrease or increase no of coins.
	// So it become necessary here to check whether the no of coins
	// to be awarded is positive or not.
	if award.Award <= 0 {
		http.Error(w, "No of coins to be awarded must be a positive number", http.StatusBadRequest)
		fmt.Println("Non-positive Award Requested. Aborted the process")
		return
	}

	user := User{Rollno: award.Rollno}
	err := validateLogin(&user, true)
	if err != nil {
		http.Error(w, "No such user", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// * The amount of _coin that a person can hold at any point
	// * in time will be capped. Thus, the total amount in the
	// * system at any point is also capped.
	if award.Award+user.Coin > MAX_COIN {
		http.Error(w, "Balance limit reached for awardee", http.StatusBadRequest)
		fmt.Println(user.Name, "is full of coin")
		return
	}
	err = awardCoin(&award)
	if err != nil {
		http.Error(w, "Unable to Award Coins", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println(award.Award, "Coins awarded to", award.Rollno)

	// TODO : Send total coins to user.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(award)
}

type transferRequest struct {
	FromRollno int `json:"from"`
	ToRollno   int `json:"to"`
	Amount     int `json:"amount"`
}

func sameBatch(firstRollno int, secondRollno int) bool {
	// Assume UG for now
	// TODO: more elaborate mechanism
	return (firstRollno/10000 == secondRollno/10000)
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	// Must be a post request
	if r.Method != "POST" {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		fmt.Println("Unsupported Method Name")
		return
	}

	if err := checkCookie(w, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// * Authorized
	var transfer transferRequest
	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		http.Error(w, "Invalid Input to the form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	//sanity checks
	Sender := User{Rollno: transfer.FromRollno}
	err = validateLogin(&Sender, true)
	if err != nil {
		fmt.Println()
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	Reciever := User{
		Rollno: transfer.ToRollno,
	}

	// It might happen that the database is busy so
	// ivalidate the request.
	err = validateLogin(&Reciever, true)
	if err != nil {
		fmt.Println()
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	if transfer.Amount <= 0 {
		http.Error(w, "No of coins to be transferred must be a positive number", http.StatusBadRequest)
		fmt.Println("Non-positive Transfer Requested. Aborted the process")
		return
	}

	if transfer.Amount > Sender.Coin {
		http.Error(w, "Not Enough Amount", http.StatusBadRequest)
		fmt.Println("Sender do not have enough money")
		return
	}

	// * The amount of _coin that a person can hold at any point
	// * in time will be capped. Thus, the total amount in the
	// * system at any point is also capped.

	// TODO: transfer Amount should be different here
	// Better check in the database.
	if transfer.Amount+Reciever.Coin > MAX_COIN {
		http.Error(w, "Balance limit reached for Reciever", http.StatusBadRequest)
		fmt.Println(Reciever.Name, "cannot recieve more coins")
		return

	}

	err = transferCoin(&transfer)
	if err != nil {
		http.Error(w, "Unable to Award Coins", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	fmt.Println(transfer.Amount, "Coins transferred from", transfer.FromRollno, "to", transfer.ToRollno)

	// TODO : Send total coins to user.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfer)
}
