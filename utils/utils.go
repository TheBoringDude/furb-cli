package utils

import (
	"net/http"
	"os"
)

// CheckInternetConnection checks if the device is connected to the internet.
// based from: https://dev.to/obnoxiousnerd/check-if-user-is-connected-to-the-internet-in-go-1hk6
func CheckInternetConnection() bool {
	_, err := http.Get("http://icanhazip.com")

	if err != nil {
		// device is offline
		return false
	}

	// device is connected online
	return true
}

// MakeDir creates dir and handles err,
// It's just that it's repetitive in the code
func MakeDir(dirname string) {
	err := os.Mkdir(dirname, 0755)
	LogErr(err)
}
