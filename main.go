package main

import (
	_ "embed"
	"os"
	"strings"

	"github.com/aixoio/idb/notice"
)

//go:embed LICENSE
var license string

func main() {
  notice.PrintLicenseNotice()
  
  if strings.Compare(strings.ToLower(os.Args[1]), "license") == 0 {
    notice.PrintLicense(license)
  }
}

