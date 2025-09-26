package cmd

import (
	"fmt"
)

var versionString string = "development"

func PrintVersion() {
	fmt.Printf("version: %v\n", versionString)
}
