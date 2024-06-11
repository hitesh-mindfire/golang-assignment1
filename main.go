package main

import (
	"assignment1/dbConfig"
	"assignment1/middlewares"
	"assignment1/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	db := dbConfig.DbConnection()
	defer db.Close()
	if db != nil {
		fmt.Println("Connected to the database!")
	} else {
		log.Fatalf("Failed to establish a connection")
	}
	router := routes.BookRouter(db)
	http.ListenAndServe(":8000", middlewares.SetJSONContentType(router))
}
