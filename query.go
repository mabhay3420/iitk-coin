// Write a program that connects to a database.
// The database that we will be working with is SQLite.
// On executing the program it should create a new table
// (say, User with two fields rollno and name).
// Create a function that takes in new user details as arguments and
// adds it to the database.
// Your program should not take any input from command line.
// Aim to write clean and structured code.

package main 

import "fmt"
import "strconv"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func AddData(ROLLNO string, NAME string) {
	database, _ := sql.Open("sqlite3", "./student.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS User (id INTEGER PRIMARY KEY, rollno TEXT, name TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO User (rollno, name) VALUES (?,?)")

	// Fill data
	statement.Exec(ROLLNO,NAME)
	// QUERY FORMAT
	// rows,_ := database.Query("SELECT id, rollno,name FROM User")
	// var id int
	// var rollno string
	// var name string

	// for rows.Next(){
	// 	rows.Scan(&id,&rollno,&name)
	// 	fmt.Println(strconv.Itoa(id) + ": " + rollno +" "+ name)
	// }

}

func main()  {
	// DUMMY DATA
	AddData("190017","Abhay Mishra")
}
