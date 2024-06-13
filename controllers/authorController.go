package controllers

import (
	"assignment1/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAuthors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM authors")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var authors []models.Author
		for rows.Next() {
			var author models.Author
			err := rows.Scan(&author.ID, &author.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			authors = append(authors, author)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(authors)
	}
}

func GetAuthorById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		log.Println(vars)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(id)
		var author models.Author
		row := db.QueryRow(`SELECT id, name FROM authors WHERE id=$1`, id)
		log.Println(row, "row")
		err = row.Scan(&author.ID, &author.Name)
		if err == sql.ErrNoRows {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(author)
	}
}

func CreateAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var author models.Author
		err := json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var id int
		err = db.QueryRow("INSERT INTO authors (name) VALUES ($1) RETURNING id", author.Name).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("id", id)
		author.ID = int(id)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Author created successfully",
			"author":  author,
		})
	}
}
