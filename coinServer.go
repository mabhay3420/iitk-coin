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

// Task

// 1. Create an endpoint that accepts a POST request and awards coins to a user.
//    The body will have the rollno of the user and the number of coins to be given.

// 2. Create an endpoint that accepts a POST request to transfer coins between two users.
//    The body will have the rollnos of the particpating users and the number of coins to transfer.

// 3. Create an endpoint that accepts a GET request and returns the coin balance of a user.
//    The body will have the roll no of the user.

// Notes

// 1. The server handles API requests concurrently by design.
//    By no way should your endpoints create or destroy coins when it is not intended.
//    You need to take care that all the steps of your transactions either complete successfully or don't happen at all.

// 2. Make sure that there is no such possible interleaving between two concurrent transactions that can cause unwanted behavior.
//    To simulate and test different interleavings you can make use of sleep timers between lines of your code.

// 3. Some of the endpoints that you will be creating would not be for all users but just admins,
//    but you can ignore that for now. We'll take up permission levels a bit later.
//    You can keep these APIs public for now i.e. no authorization required.

// 4. Take care of as many edge cases as you can.

// Possible edge cases
// Award:
// 1. negative or 0 amount : need to check
// 2. Does not fit in range of integer : json decode should throw an error ideally : need to check
// 3. User does not exist : authorization

// transfer
// 1. Invalid amount : 0 or negative or sender do not enough amount : need to check
// 2. Sender Unauthorized : authorization
// 2. Second user does not exist : error while getting user data.

//
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

	user := User{ Rollno: balance.Rollno}
	err = validateLogin(&user,true) // has a cookie
	if err != nil {
		// extra sanity check
		fmt.Println("Invalid id/passoword")
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}
	response := Response{ Rollno: user.Rollno, Name: user.Name, Coin: user.Coin}

	fmt.Println("Balance of",user.Name,"is",user.Coin)
	w.Header().Set("Content-Type","application/json")
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
	Sender := User{
		Rollno: transfer.FromRollno,
	}
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
