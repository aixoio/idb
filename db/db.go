package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() (*sqlx.DB, error) {
  db, err := sqlx.Open("sqlite3", "./idb.sqlite3")
  if err != nil {
    return nil, err
  }

  return db, db.Ping()
}

func CloseDB(db *sqlx.DB) {
  db.Close()
}

