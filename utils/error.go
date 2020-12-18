package utils

import (
	"fmt"
	"os"
)

// ExitErr show error message and stop the app,
func ExitErr(message string) {
	fmt.Println("\n " + message)
	os.Exit(1) // stop the app
}