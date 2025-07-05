package database

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func GetDbConnection() (*sql.DB, error) {

	//TODO Add support for local development

	url := ""

	db, err := sql.Open("libsql", url)
	return db, err
}
