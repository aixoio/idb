package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/aixoio/idb/db"
	"github.com/aixoio/idb/help"
	"github.com/aixoio/idb/notice"
)

//go:embed LICENSE
var license string

func main() {
  notice.PrintLicenseNotice()
  
  if len(os.Args) <= 1 {
    help.PrintOnelineHelp()
    return
  }

  if strings.Compare(strings.ToLower(os.Args[1]), "help") == 0 {
    help.PrintFullHelp()
    return
  }

  if strings.Compare(strings.ToLower(os.Args[1]), "license") == 0 {
    notice.PrintLicense(license)
    return
  }

  d, err := db.ConnectToDB()
  if err != nil {
    panic(err)
  }
  defer db.CloseDB(d)
  
  err = db.CreateDatabaseSchemaIfNotExists(d)
  if err != nil {
    panic(err)
  }

  if strings.Compare(strings.ToLower(os.Args[1]), "add") == 0 && len(os.Args) >= 3 {
    err := db.NewIdea(d, os.Args[2])
    if err != nil {
      panic(err)
    }
    
    fmt.Println("Added your idea to the database")
    return
  }
}

