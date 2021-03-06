package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* GLOBAL VARIABLES*/
var MAX_COIN int = 100000        // max no of coins a person can hold
var MIN_ACTIVE int = 5           // minimum no of activites student must be involved in order to transfer/redeem money
var ADMIN = []int{190058}        // ADMIN list, used when assigning role
var COUNCIL_CORE = []int{190345} // COUNCIL CORE members
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

	rollno, err := checkCookie(w, r)
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

	// The user logged in can see only their
	// balance. Admin is exceptional.
	if (user.Rollno != rollno) && (user.Role != "ADMIN") {
		http.Error(w, "Unauthorized", http.StatusBadRequest)
		fmt.Println(rollno, "tried to see the balance of", user.Rollno, "Denied!")
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
	rollno, err := checkCookie(w, r)
	if err != nil {
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

	// sanity checks
	if award.Award <= 0 {
		http.Error(w, "No of coins to be awarded must be a positive number", http.StatusBadRequest)
		fmt.Println("Non-positive Award Requested. Aborted the process")
		return
	}
	curr_user := User{Rollno: rollno}
	err = validateLogin(&curr_user, true)
	if err != nil {
		http.Error(w, "No such user", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	user := User{Rollno: award.Rollno}
	err = validateLogin(&user, true)
	if err != nil {
		http.Error(w, "No such user", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// ONLY ADMINS Allowed to acces this endpoint
	if curr_user.Role != "ADMIN" {
		http.Error(w, "You need admin role to award others", http.StatusBadRequest)
		fmt.Println(curr_user.Rollno, "tried to award", user.Rollno)
		return
	}

	// ADMINS cannot reward themselves
	if user.Role == "ADMIN" {
		http.Error(w, "No of coins in admin account cannot change", http.StatusBadRequest)
		fmt.Println(user.Rollno, "who has admin access tried to award himself.Aborted!")
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

	rollno, err := checkCookie(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// ! Unnecessary lookup : find some better method later.
	curr_user := User{Rollno: rollno}

	err = validateLogin(&curr_user, true)
	if err != nil {
		fmt.Println()
		http.Error(w, "No such user", http.StatusBadRequest)
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

	Reciever := User{Rollno: transfer.ToRollno}

	// It might happen that the database is busy so
	// ivalidate the request.
	err = validateLogin(&Reciever, true)
	if err != nil {
		fmt.Println()
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	// CHECK PERMISSION
	// Most important : Sender must be the one who is logged in
	if Sender.Rollno != curr_user.Rollno {
		http.Error(w, "You can transfer from your account only", http.StatusBadRequest)
		fmt.Println(curr_user.Rollno, "tried to send from", Sender.Rollno, "Account. Aborted!")
		return
	}

	// Sender must participate in more than MIN_ACTIVE activites
	if Sender.Activity < MIN_ACTIVE {
		http.Error(w, "You are not eligible to transfer coins, Participate More", http.StatusBadRequest)
		fmt.Println(Sender.Rollno, "Has participated in only", Sender.Activity, "Activities. Unable to transfer!")
		return
	}

	// 1. ADMIN : not allowed in transfers
	if (Reciever.Role == "ADMIN") || (Sender.Role == "ADMIN") {
		http.Error(w, "Balance of Admin cannot change in any way", http.StatusBadRequest)
		fmt.Println("Admin was involved in transfers. Aborted.")
		return
	}

	// 2. COUNCIL_CORE : NOT allowed in transfers,
	// though can redeem coins and can be awarded by admins
	if (Sender.Role == "COUNCIL_CORE") || (Reciever.Role == "COUNCIL_CORE") {
		http.Error(w, "Council core members can only spend the money or get award from admin", http.StatusBadRequest)
		fmt.Println("Council core member was involved in transfers. Aborted.")
	}

	// Now Check amounts
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

	// The amount of _coin that a person can hold at any point
	// in time will be capped. Thus, the total amount in the
	// system at any point is also capped.

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
