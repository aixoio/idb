package db

import (
	_ "embed"

	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var schema string

func CreateDatabaseSchemaIfNotExists(db *sqlx.DB) error { 
  _, err := db.Exec(schema)
  return err
}

