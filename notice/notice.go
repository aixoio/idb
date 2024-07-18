package notice

import "fmt"

func PrintLicenseNotice() {
  fmt.Println("This program is licensed under the MIT license")
  fmt.Println("Run: 'idb license' for a copy of the MIT license")
}

func PrintLicense(license string) {
  fmt.Println(license)
}

