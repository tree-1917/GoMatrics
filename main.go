package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"gomatrics/router"
)

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/gomatric")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Give DB to router package
	router.DB = db

	r := mux.NewRouter()

	// Register routes
	router.RegisterPostRoutes(r)

	fmt.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
