package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		fmt.Printf("Available Drivers are: %v\n\n", driver)
	}
}

func openDatabase() (db *sql.DB, err error) {
	if db, err = sql.Open("sqlite", "products.db"); err != nil {
		log.Println("Could not open database connection.")
	} else {
		fmt.Printf("Opened database: %T\n", db)
	}
	return
}

func openPGDatabase() {
	// "postgres://<user-name>:<password>@<db-hostname>/<db-name>"
	connStr := "postgres://postgres:@localhost/go_development"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("Opened database: %T\n", db)
	}
}
