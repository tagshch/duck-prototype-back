package storages

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteStorage struct {
	DB *sql.DB
}

func NewSqlLiteStorage(file string) (*SqlLiteStorage, error) {
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		return nil, err
	}

	return &SqlLiteStorage{
		DB: db,
	}, nil
}
