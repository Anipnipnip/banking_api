package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=banking_db sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
