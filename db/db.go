package db

import (
  "github.com/jmoiron/sqlx"
  _ "github.com/mattn/go-sqlite3"
)

func ConnectToDB() (*sqlx.DB, error) {
  return sqlx.Open("sqlite3", "./idb.sqlite3")
}

