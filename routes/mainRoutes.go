package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func MainRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	BookRouter(db, r)
	AuthorRouter(db, r)
	return r
}
