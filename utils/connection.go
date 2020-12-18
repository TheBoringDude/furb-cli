package utils

import "net/http"

// CheckInternetConnection checks if the device is connected to the internet.
// based from: https://dev.to/obnoxiousnerd/check-if-user-is-connected-to-the-internet-in-go-1hk6
func CheckInternetConnection() bool{
	_, err := http.Get("http://icanhazip.com")

	if err != nil {
		// device is offline
		return false
	}

	// device is connected online
	return true
}