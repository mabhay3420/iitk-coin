// 1. Create web server with two endpoints /login and /signup that accepts POST requests.

// 2. The /signup endpoint will receive a new user's roll no,
// password and other details and create a new user in the database.
//  You already have created a function earlier to add users to the database.
//   Don't store the password as plain text. Apply hashing and salting.

// 3. The /login endpoint will take in the Rollno and password and
// if verified successfully will return a JWT (JSON Web Token) as part of the response.

// 4. Create an endpoint, say /secretpage that returns
// some dummy data only if the user is logged in. By logged in here we mean
// that the JWT sent along the request should be a valid token and the user
// is authorized to access the endpoint.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "strconv"
)

type Response struct{
	Rollno int `json:"rollno"`
	Name string `json:"name"`
	Coin int `json:"coin"`
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
	err = getUserInfo(&user)

	if err != nil {
		http.Error(w, "invalid user rollno/password combination", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	response := Response{user.Rollno,user.Name,user.Coin}
	w.Header().Set("Content-Type","application/json")
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
	response := Response{user.Rollno,user.Name,user.Coin}
	w.Header().Set("Content-Type","application/json")
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

func startServer() {

	// Handle incoming requests
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)

	// Start the server
	fmt.Println("Starting Server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Listen and server error:", err)
	}
}
