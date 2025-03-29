package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
)

func getDBConn() *sql.DB {
	// this (sql.DB) is supposedely thread safe but I am having trouble letting other goroutines
	// use it (nil pointer exception or the like)
	db, err := sql.Open("sqlite", "file:/home/dev1/data/bb/pass.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	fmt.Println("starting...")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleLogin)
	err := http.ListenAndServe(":3000", mux)
	if nil != err {
		log.Fatal(err)
	}
}
