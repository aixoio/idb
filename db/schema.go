package db

import (
	_ "embed"
	"time"

	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var schema string

type Idea struct {
  Id int
  Idea string
  Created_At time.Time
  Updated_At time.Time
}

func CreateDatabaseSchemaIfNotExists(db *sqlx.DB) error { 
  _, err := db.Exec(schema)
  return err
}

