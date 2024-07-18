package db

import "github.com/jmoiron/sqlx"

func NewIdea(db *sqlx.DB, idea string) error {
  _, err := db.Exec(db.Rebind("INSERT INTO ideas (idea) VALUES (?)"), idea)
  return err
}

func GetAll(db *sqlx.DB) ([]Idea, error) {
  ideas := []Idea{}
  return ideas, db.Select(&ideas, "SELECT * FROM ideas ORDER BY id ASC")
}

