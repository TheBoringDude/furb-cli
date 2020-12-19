package furb

import (
	"fmt"
	"net/url"
	"os"

	"github.com/TheBoringDude/furb-cli/utils"
)

// Furb main instance session
type Furb struct {
	Request string
	Type    string
}

// InitConf => checks connection and argument site
func (f *Furb) InitConf() {
	// check internet connection
	onl := utils.CheckInternetConnection()
	if !onl {
		fmt.Println("\n [!] NOTE: You are not connected to the internet. Please connect and try again.")
		os.Exit(1)
	}

	// check if website arg is valid or not
	_, err := url.ParseRequestURI(f.Request)
	utils.ExitErr(err, "[!] NOTE: Manga url is not valid!")
}
