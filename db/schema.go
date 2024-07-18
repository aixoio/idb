package db

import (
	_ "embed"
	"fmt"
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

func (i Idea) String() string {
  return fmt.Sprintf("%d | %s | %s | %s", i.Id, i.Idea, i.Created_At.Local().String(), i.Updated_At.Local().String())
}

func CreateDatabaseSchemaIfNotExists(db *sqlx.DB) error { 
  _, err := db.Exec(schema)
  return err
}

