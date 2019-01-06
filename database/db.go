package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dsn string) (*sql.DB, error) {
	DB, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(1)

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
