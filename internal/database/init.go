package database

import (
	"database/sql"
	"log"
)

func InitDatabase(dbUrl string) *Queries {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := New(db)

	return dbQueries
}
