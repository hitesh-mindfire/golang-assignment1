package dbConfig

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
