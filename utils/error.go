package utils

import (
	"fmt"
	"os"
	"log"
)

// ExitErr show error message and stop the app,
func ExitErr(err error, message string) {
	if err != nil {
		fmt.Println("\n " + message)
		os.Exit(1) // stop the app
	}
}

// LogErr logs the error
func LogErr(err error){
	if err != nil {
		log.Fatalln(err)
	}
}