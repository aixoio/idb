package db

import "github.com/jmoiron/sqlx"

func NewIdea(db *sqlx.DB, idea string) error {
  _, err := db.Exec(db.Rebind("INSERT INTO ideas (idea) VALUES (?)"), idea)
  return err
}

