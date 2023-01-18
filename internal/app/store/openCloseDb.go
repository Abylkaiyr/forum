package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func (s *Store) Open() error {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
