package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite", ".url-shortener.db")
	if err != nil {
		return nil, fmt.Errorf("%s : %m", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS url (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    alias TEXT NOT NULL UNIQUE,
	    url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	)`)
	if err != nil {
		return nil, fmt.Errorf("%s : %m", op, err)
	}
}