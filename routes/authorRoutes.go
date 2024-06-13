package routes

import (
	"assignment1/controllers"
	"assignment1/utils"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthorRouter(db *sql.DB, router *mux.Router) {
	router.HandleFunc("/authors", controllers.GetAuthors(db)).Methods("GET")
	router.HandleFunc("/authors/{id}", controllers.GetAuthorById(db)).Methods("GET")
	router.HandleFunc("/authors", controllers.CreateAuthor(db)).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(utils.NotFoundHandler())
}
