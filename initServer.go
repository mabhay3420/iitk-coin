package main

import (
	"encoding/json"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	// "time"
	// "strconv"
)

type dummyResponse struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

func secretHandler(w http.ResponseWriter, r *http.Request) {

	if err := checkCookie(w, r); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	dummyData := dummyResponse{"We are so excited to meet you", "invictus"}
	// Status OK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummyData)

	// success
	fmt.Println("Secret content Delievered to someone") // modify later to get name
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// TODO : Learn More about Methods
	// Valid Methods
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
		fmt.Println("Invlid Method")
		// Error method does not necessarly close the request. Need to return.
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hii there, Move to either login or signup")
}

func startServer() {

	// Handle incoming requests
	http.HandleFunc("/", homeHandler)             // anyone
	http.HandleFunc("/login", loginHandler)       // anyone
	http.HandleFunc("/signup", signupHandler)     // anyone
	http.HandleFunc("/secret", secretHandler)     // anyone
	http.HandleFunc("/award", awardHandler)       // admin only
	http.HandleFunc("/transfer", transferHandler) // only sender can request
	http.HandleFunc("/balance", balanceHandler)   // anyone

	// Start the server
	fmt.Println("Starting Server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Listen and server error:", err)
	}
}
