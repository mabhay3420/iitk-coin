// 1. Create web server with two endpoints /login and /signup that accepts POST requests.

// 2. The /signup endpoint will receive a new user's roll no,
// password and other details and create a new user in the database.
//  You already have created a function earlier to add users to the database.
//   Don't store the password as plain text. Apply hashing and salting.

// 3. The /login endpoint will take in the rollno and password and
// if verified successfully will return a JWT (JSON Web Token) as part of the response.

// 4. Create an endpoint, say /secretpage that returns
// some dummy data only if the user is logged in. By logged in here we mean
// that the JWT sent along the request should be a valid token and the user
// is authorized to access the endpoint.

package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// // Correct Endpoint
	// if r.URL.Path != "/login" {
	// 	http.Error(w, "404 Not found.", http.StatusNotFound)
	// 	return
	// }

	// TODO : Learn More about Methods
	// Valid Methods
	if (r.Method != "POST") && (r.Method != "GET") {
		http.Error(w, "Method is not Supported.", http.StatusNotFound)
		return
	}


	// Do Something
	fmt.Fprint(w, "Hello This is login Endpoint\n")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Correct Endpoint
	if r.URL.Path != "/signup" {
		http.Error(w, "404 Not found.", http.StatusNotFound)
		fmt.Println("Invalid Endpoint")
		return
	}

	// TODO : Learn More about Methods
	// Valid Methods
	// if (r.Method != "POST" && r.Method != "GET") {
	// 	http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
	// 	fmt.Println("Invlid Method")
	// 	// Error method does not necessarly close the request. Need to return.
	// 	return
	// }
	// Don't Know what to do.
	if(r.Method == "GET"){
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello This is signup Endpoint\n")
		return
	}

	// POST Format: rollno:-- name:--

	if err := r.ParseForm(); err != nil{
		http.Error(w,"Invalid input to Form",http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// get fields from request
	rollno := r.FormValue("rollno")
	name := r.FormValue("name")
	password := r.FormValue("password")

	// testing purpose

	// Add to the user Database
	// TODO: Find better method to store as int
	roll_int,err := strconv.Atoi(rollno)

	//? Shouldn't we use Errors which signify what went wrong.
	if err != nil{
		http.Error(w,"Roll number must be integer\n",http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if(password == "" || name==""){
		http.Error(w, "User Name/Password Can not be empty", http.StatusBadRequest)
		fmt.Println("User Name/Password empty")
		return
	}
	// Adding new user
	err = addUser(roll_int,name,password)

	if(err != nil){
		http.Error(w, "Invalid User Name/Password or User Already Exists", http.StatusBadRequest)
		fmt.Println("Invalid Request, DB error:",err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w,"POST Request Succesful.\nUser Succesfully Added\nYour Rollno:%s Your Name:%s\n",rollno,name)

	// ! Should not write to header twice: http: superfluous response.WriteHeader call
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	// Correct Endpoint
	// if r.URL.Path != "/signup" {
	// 	http.Error(w, "404 Not found.", http.StatusNotFound)
	// 	fmt.Println("Invalid Endpoint")
	// 	return
	// }

	// TODO : Learn More about Methods
	// Valid Methods
	if ( r.Method != "GET") {
		http.Error(w, "Method is not Supported.", http.StatusMethodNotAllowed)
		fmt.Println("Invlid Method")
		// Error method does not necessarly close the request. Need to return.
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w,"Hii there, Move to either login or signup")
}

func startServer() {

	// Handle incoming requests
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)

	// Start the server
	fmt.Println("Starting Server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("Listen and server error:", err)
	}
}
