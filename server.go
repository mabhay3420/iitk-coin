// TODO : A refresh endpoint.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
	// "strconv"
)

// ? What information should we send to the client side?
// ? Just the status code or some information in statement form?
// ? If the later what type of information are sensitive and
// ? should not be exposed to the client?

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
	err = validateLogin(&user)

	if err != nil {
		http.Error(w, "invalid user rollno/password combination", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}

	// Valid User

	// Step 1: Creating a JWT for User.

	// Step 1a) Create a Payload for JWT.

	// Expiration time of token : 1 min for now : need to refresh
	expirationTime := time.Now().Add(1 * time.Minute)
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
	// Verify User
	// TODO : Create a function which does this authorization
	// 	  so that we can use it frequently.
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
	if award.Award <= 0{
		http.Error(w,"No of coins to be awarded must be a positive number",http.StatusBadRequest)
		fmt.Println("Non-positive Award Requested. Aborted the process")
		return
	}
	
	err := updateUserCoin(award)
}

func startServer() {

	// Handle incoming requests
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/secret", secretHandler)

	// Start the server
	fmt.Println("Starting Server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Listen and server error:", err)
	}
}
