package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type Response struct {
	Rollno int    `json:"rollno"`
	Name   string `json:"name"`
	Coin   int    `json:"coin"`
}

var jwtKey = []byte("hello_world")

// read request
type Credentials struct {
	Rollno   int    `json:"rollno"`
	Password string `json:"password"`
}

// StandardClaims will be used to provide fields like expiry time
type Claims struct {
	Rollno int `json:"rollno"`
	jwt.StandardClaims
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.WriteHeader(200)
		fmt.Fprint(w, "hello this is login endpoint\n")
		return
	}

	// Verify user
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input to the form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	err = validateLogin(&user, false)

	if err != nil {
		http.Error(w, "invalid user rollno/password combination", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// Issue new cookie every time
	// Step 1: Creating a JWT for User.
	// Step 1a) Create a Payload for JWT.
	// Expiration time of token : 5 min for now
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Rollno: user.Rollno,
		StandardClaims: jwt.StandardClaims{
			// Must be in unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Step 1b) Declare a token with algo used for hashing and the claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Step 1c) Create the JWT string by hashing with key.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// Error in creating the JWT
		http.Error(w, "Error in creating JWT", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Step 2: Set Client Cookie for "token"
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: expirationTime,
	})

	// Step 3: Usual Information
	response := Response{user.Rollno, user.Name, user.Coin}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Println("Login of", user.Name, "Succesful")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.WriteHeader(200)
		fmt.Fprint(w, "hello this is signup endpoint\n")
		return
	}

	var user User
	// get fields from request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input to form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Adding new user : Rollno, Name , Password are expected.
	err = addUser(&user)

	if err != nil {
		http.Error(w, "invalid user name/password or user already exists", http.StatusBadRequest)
		fmt.Println("invalid Request, DB error:", err)
		return
	}
	response := Response{user.Rollno, user.Name, user.Coin}
	w.Header().Set("Content-Type", "application/json")
	// status OK
	json.NewEncoder(w).Encode(response)
}

func checkCookie(w http.ResponseWriter, r *http.Request) (int,error) {
	// Verify User
	token, err := r.Cookie("jwt")

	// Error in Cookie extraction
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie means user is not logged in.
			return 0,err
		}

		// Other type of Errors
		return 0,err
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
			return 0,err
		}

		// Unknown Error
		return 0,err
	}

	// Expired
	if !tkn.Valid {
		return 0,fmt.Errorf("invalid token")
	}

	return claims.Rollno,nil
}
