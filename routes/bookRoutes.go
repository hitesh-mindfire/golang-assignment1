package routes

import (
	"assignment1/controllers"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 - Not Found", http.StatusNotFound)
}

func BookRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID(db)).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	// router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.GetBooks(w, r, db)
	// }).Methods("GET")   // one way
	return router
}
