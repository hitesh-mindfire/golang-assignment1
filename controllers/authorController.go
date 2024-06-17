package controllers

import (
	"assignment1/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get all authors
// @Description Retrieves a list of all authors.
// @Produce json
// @Success 200 {array} models.Author
// @Router /authors [get]
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

// @Summary Get an author by ID
// @Description Retrieves an author based on provided ID.
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {object} models.Author
// @Router /authors/{id} [get]
func GetAuthorById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var author models.Author
		row := db.QueryRow(`SELECT id, name FROM authors WHERE id=$1`, id)
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

// @Summary Create a new author
// @Description Creates a new author with provided data.
// @Accept json
// @Produce json
// @Param author body models.Author true "Author object"
// @Success 201 {object} models.Author
// @Router /authors [post]
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
		author.ID = int(id)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Author created successfully",
			"author":  author,
		})
	}
}

// @Summary Update an existing author
// @Description Updates an existing author with provided data.
// @Accept json
// @Produce json
// @Param id path string true "Author ID"
// @Param author body models.Author true "Author object"
// @Success 200 {object} models.Author
// @Router /authors/{id} [put]
func UpdateAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM authors WHERE id = $1", id).Scan(&count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count == 0 {
			http.Error(w, "No record found", http.StatusNotFound)
			return
		}

		var author models.Author
		err = json.NewDecoder(r.Body).Decode(&author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE authors SET name=$1 WHERE id=$2", author.Name, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		author.ID = id
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Author updated successfully",
			"author":  author,
		})
	}
}

// @Summary Delete an author
// @Description Deletes an author based on provided ID.
// @Param id path string true "Author ID"
// @Success 204 "No Content"
// @Router /authors/{id} [delete]
func DeleteAuthor(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM authors WHERE id = $1", id).Scan(&count)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count == 0 {
			http.Error(w, "No record found", http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM authors WHERE id=$1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":   "Author deleted successfully",
			"author_id": id,
		})
	}
}
