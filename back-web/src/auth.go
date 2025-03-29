package main

import (
	"fmt"
	//"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
)

//err = bcrypt.CompareHashAndPassword(db_hashed, password)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := "cubesnail"

	db := getDBConn()
	defer db.Close()

	result := db.QueryRow(`select username, password from users where username=?`, username)
	var db_username string
	var db_hashed []byte
	err := result.Scan(&db_username, &db_hashed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DEBUG(user:%s hashed:%s)\n", db_username, db_hashed)

	io.WriteString(w, db_username)
}
