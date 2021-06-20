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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// // Correct Endpoint
	// if r.URL.Path != "/login" {
	// 	http.Error(w, "404 Not found.", http.StatusNotFound)
	// 	return
	// }

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
		fmt.Fprint(w, "Hello This is login Endpoint\n")
		return
	}

	// Verify user
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Input to the form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	// user.Rollno, err = strconv.Atoi(r.FormValue("Rollno"))
	// user.password = r.FormValue("password")

	// //? Shouldn't we use Errors which signify what went wrong.
	// if err != nil {
	// 	http.Error(w, "Roll number must be integer\n", http.StatusBadRequest)
	// 	fmt.Println(err)
	// 	return
	// }

	// Should Be handled on database side
	// if(user.password==""){
	// 	http.Error(w,"Password Cannot be empty",http.StatusBadRequest)
	// 	fmt.Println("Empty password")
	// 	return
	// }
	err = getUserInfo(&user)

	if err != nil {
		http.Error(w, "Invalid User Rollno/Password Combination", http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type","application/json")
	// Status OK
	json.NewEncoder(w).Encode(user)
	fmt.Println("Login of", user.Name, "Succesful")

	// Compare with database Password.

}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Correct Endpoint
	// if r.URL.Path != "/signup" {
	// 	http.Error(w, "404 Not found.", http.StatusNotFound)
	// 	fmt.Println("Invalid Endpoint")
	// 	return
	// }

	// TODO : Learn More about Methods
	// Valid Methods
	// if (r.Method != "POST" && r.Method != "GET") {
	// 	http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
	// 	fmt.Println("Invlid Method")
	// 	// Error method does not necessarly close the request. Need to return.
	// 	return
	// }
	// Don't Know what to do.
	if r.Method == "GET" {
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello This is signup Endpoint\n")
		return
	}

	// POST Format: Rollno:-- Name:--

	var user User
	// get fields from request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input to Form", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	user.Coin = 0
	// user.Rollno, err := strconv.Atoi(r.FormValue("Rollno"))
	// if err != nil {
	// 	http.Error(w, "Invalid input to Form", http.StatusBadRequest)
	// 	fmt.Println(err)
	// 	return
	// }
	// user := User{
	// 	Rollno:   roll_int,
	// 	Name:     r.FormValue("Name"),
	// 	password: r.FormValue("password"),
	// 	Coin:     0,
	// }

	// // Add to the user Database
	// // TODO: Find better method to store as int
	// //? Shouldn't we use Errors which signify what went wrong.
	// if err != nil {
	// 	http.Error(w, "Roll number must be integer\n", http.StatusBadRequest)
	// 	fmt.Println(err)
	// 	return
	// }

	// Will be performed on
	// if password == "" || Name == "" {
	// 	http.Error(w, "User Name/Password Can not be empty", http.StatusBadRequest)
	// 	fmt.Println("User Name/Password empty")
	// 	return
	// }
	// Adding new user
	err = addUser(&user)

	if err != nil {
		http.Error(w, "Invalid User Name/Password or User Already Exists", http.StatusBadRequest)
		fmt.Println("Invalid Request, DB error:", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "POST Request Succesful.\nUser Succesfully Added\nYour Rollno:%d Your Name:%s\n", user.Rollno, user.Name)

	// ! Should not write to header twice: http: superfluous response.WriteHeader call
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Correct Endpoint
	// if r.URL.Path != "/signup" {
	// 	http.Error(w, "404 Not found.", http.StatusNotFound)
	// 	fmt.Println("Invalid Endpoint")
	// 	return
	// }

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
