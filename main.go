package main

import (
	"assignment1/dbConfig"
	_ "assignment1/docs"
	"assignment1/middlewares"
	"assignment1/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// @title Golang Assignment
// @version 1.0
// @description This is a golang assignment1
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

	router := routes.MainRouter(db)
	router.Use(middlewares.SetJSONContentType)
	log.Fatal(http.ListenAndServe(":8000", router))
}
