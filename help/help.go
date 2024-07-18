package help

import "fmt"

func PrintOnelineHelp() {
  fmt.Println("Usage: idb <command> [input]")
}

func PrintFullHelp() {
  fmt.Println("Help menu:")
  fmt.Println("\tidb license: Prints license infomation")
}

