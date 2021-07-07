package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	// "time"
	// "strconv"
)

type dummyResponse struct {
	Message string `json:"message"`
	User    string `json:"user"`
	From    string `json:"from"`
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("jwt")

	// Error in Cookie extraction
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie means user is not logged in.
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		// Other type of Errors
		http.Error(w, "Unauthorized!", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// token present

	tokenString := token.Value

	// New instance of claims
	claims := &Claims{}

	// Parse the JWT string and store in `claims`.
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Problem with token
	if err != nil {
		// Signature did not match.
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		// Unknown Error
		http.Error(w, "Unauthorized", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Expired
	// ? Invalid/Expired claims will already throw an error in previous Step
	// ? Then how can this token become invalid?
	if !tkn.Valid {
		http.Error(w, "You need to log in Again!", http.StatusUnauthorized)
		fmt.Println("Token Expired")
		return
	}

	dummyData := dummyResponse{"We are so excited to meet you", claims.Name, "invictus"}
	// Status OK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummyData)

	// success
	fmt.Println("Secret content Delievered to ", claims.Name)
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
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/secret", secretHandler)
	http.HandleFunc("/award", awardHandler)
	http.HandleFunc("/transfer", transferHandler)
	http.HandleFunc("/balance",balanceHandler)

	// Start the server
	fmt.Println("Starting Server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Listen and server error:", err)
	}
}
