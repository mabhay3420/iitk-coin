// TODO : A refresh endpoint.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	// "log"
	"net/http"
	"time"
	// "strconv"
)

type Response struct {
	Rollno int    `json:"rollno"`
	Name   string `json:"name"`
	Coin   int    `json:"coin"`
}

// ? Find some secure method to save the key
//https://www.sohamkamani.com/golang/jwt-authentication/#the-jwt-format

// JWT key used to create the signature
var jwtKey = []byte("hello_world")

// read request
type Credentials struct {
	Rollno   int    `json:"rollno"`
	Password string `json:"password"`
}

// Struct Which will be encoded to a JWT
// StandardClaims will be used to provide fields like expiry time

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : Learn More about Methods
	// Valid Methods
	// if (r.Method != "POST" && r.Method != "GET") {
	// 	http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
	// 	fmt.Println("Invlid Method")
	// 	// Error method does not necessarly close the request. Need to return.
	// 	return
	// }

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

	// issue new cookie every time
	// Valid User

	// Step 1: Creating a JWT for User.

	// Step 1a) Create a Payload for JWT.

	// Expiration time of token : 5 min for now : need to refresh
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			// Must be in unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Step 1b) Declare a token with algo used for hashing and the claims.

	// TODO: Learn a bit about HS256.
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
	// Status OK
	json.NewEncoder(w).Encode(response)
	fmt.Println("Login of", user.Name, "Succesful")

}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : Learn More about Methods
	// Valid Methods
	// if (r.Method != "POST" && r.Method != "GET") {
	// 	http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
	// 	fmt.Println("Invlid Method")
	// 	// Error method does not necessarly close the request. Need to return.
	// 	return
	// }

	if r.Method == "GET" {
		w.WriteHeader(200)
		fmt.Fprint(w, "hello this is signup endpoint\n")
		return
	}

	// POST format Rollno

	var user User
	// get fields from request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input to form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	user.Coin = 0
	// Adding new user
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

func checkCookie(w http.ResponseWriter, r *http.Request) error {
	// Verify User
	token, err := r.Cookie("jwt")

	// Error in Cookie extraction
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie means user is not logged in.
			// http.Error(w, "No cookie faun", http.StatusUnauthorized)
			// fmt.Println(err)
			return err
		}

		// Other type of Errors
		// http.Error(w, "Unauthorized!", http.StatusBadRequest)
		// fmt.Println(err)
		return err
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
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			// fmt.Println(err)
			return err
		}

		// Unknown Error
		// http.Error(w, "Unauthorized", http.StatusBadRequest)
		// fmt.Println(err)
		return err
	}

	// Expired
	// ? Invalid/Expired claims will already throw an error in previous Step
	// ? Then how can this token become invalid?
	if !tkn.Valid {
		// http.Error(w, "you need to log in again!", http.StatusUnauthorized)
		// fmt.Println("Token Expired")
		return fmt.Errorf("invalid token")
	}

	return nil

}
