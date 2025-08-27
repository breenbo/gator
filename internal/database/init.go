package database

import (
	"database/sql"
	"fmt"
	"os"
)

func InitDatabase(dbUrl string) *Queries {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := New(db)

	return dbQueries
}
