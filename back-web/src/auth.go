package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type Req struct {
	Username string
	Password string
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var parsed_body Req
	err := decoder.Decode(&parsed_body)
	if err != nil {
		http.Error(w, "couldnt parse username/password from body", http.StatusBadRequest)
		return
	}

	db := getDBConn()
	defer db.Close()

	result := db.QueryRow(`select password from users where username=?`, parsed_body.Username)

	var db_hashed []byte
	err = result.Scan(&db_hashed)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "login unsuccessful (user not found)", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Fatal(err)
	} else {
		err = bcrypt.CompareHashAndPassword(db_hashed, []byte(parsed_body.Password))
		if err != nil {
			http.Error(w, "login unsuccessful (incorrect password)", http.StatusUnauthorized)
			return
		} else {
			io.WriteString(w, "login successful!")
		}
	}
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var parsed_body Req
	err := decoder.Decode(&parsed_body)
	if err != nil {
		http.Error(w, "couldnt parse username/password from body", http.StatusBadRequest)
		return
	}

	db := getDBConn()
	defer db.Close()

	var hashed []byte
	res, err := bcrypt.GenerateFromPassword([]byte(parsed_body.Password), bcrypt.DefaultCost)
	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			http.Error(w, "password too long", http.StatusBadRequest)
			return
		} else {
			log.Fatal(err)
		}
	}
	hashed = res

	_, err = db.Exec(`insert into users (username, password) values(?, ?)`,
		parsed_body.Username, hashed)

	if err != nil {
		if liteErr, ok := err.(*sqlite.Error); ok && liteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		} else {
			log.Fatal(err)
		}
	} else {
		io.WriteString(w, "user created!")
	}
}
