package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "file:/home/dev1/data/bb/pass.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("create table if not exists users(id integer primary key, username text, password blob)")
	if err != nil {
		log.Fatal(err)
	}

	var choice int
	fmt.Print("1. register 2. login >> ")
	fmt.Scanf("%v", &choice)

	switch choice {
	case 1:
		fmt.Println("register new user")

		var username string
		var password []byte
		fmt.Print("username: ")
		fmt.Scanf("%v", &username)
		fmt.Print("password: ")
		fmt.Scanf("%v", &password)

		hashed, _ := bcrypt.GenerateFromPassword(password, 10)

		query := `insert into users (username, password) values (?, ?)`
		_, err = db.Exec(query, username, hashed)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s : %s -> %s\n", username, password, hashed)
	case 2:
		fmt.Println("login")

		var username string
		var password []byte
		fmt.Print("username: ")
		fmt.Scanf("%v", &username)
		fmt.Print("password: ")
		fmt.Scanf("%v", &password)

		result := db.QueryRow(`select username, password from users where username=?`, username)
		var db_username string
		var db_hashed []byte
		err = result.Scan(&db_username, &db_hashed)
		fmt.Printf("(%s %s)\n", db_username, db_hashed)

		err := bcrypt.CompareHashAndPassword(db_hashed, password)
		if err != nil {
			fmt.Printf("%s login failed", username)
		} else {
			fmt.Printf("%s login successful", username)
		}
	default:
		fmt.Println("huh?")
	}

}
