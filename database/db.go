package database

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func GetDbConnection() (*sql.DB, error) {

	url := ""

	db, err := sql.Open("libsql", url)

	return db, err
}
