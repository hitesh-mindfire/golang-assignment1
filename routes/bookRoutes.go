package routes

import (
	"assignment1/controllers"
	"assignment1/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func BookRouter(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById(db)).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler())
	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.GetBooks(w, r, db)
	// }).Methods("GET")   // one way
}
