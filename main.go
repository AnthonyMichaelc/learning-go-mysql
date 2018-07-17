package main

// importing needed tools
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// create a struct for pulling name in json
	Name string `json:"name"`
}

func main() {
	fmt.Println("Go MySql")

	// configuration/authenticating into the db
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/testdb")

	if err != nil { // display errors connecting to localhost db
		panic(err.Error())
	}
	defer db.Close()

	// Inserting into the mysql

	// insert, err := db.Query("INSERT INTO users VALUES('Ant')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()
	// fmt.Println("successful connection to mysql  db")
	// fmt.Println("successful insert to users table")

	// Reading from the db/display in terminal
	results, err := db.Query("SELECT name FROM users")
	if err != nil { // displaying errors from executing sql command
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		if err != nil { // displaying errors getting the name field
			panic(err.Error())
		}

		fmt.Println("My name is", user.Name) // printing out my final result
	}
}
