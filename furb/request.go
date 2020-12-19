package furb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// main api url vars
var apiURL = "https://magna-sc.cf" // do not add `/` in here
var mangaAPI = "/manga"
var chapterAPI = "/chapters"
var q = "?q="

// ReqAPI => requests a JSON response from an API.
// Based from: https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7 & https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
func (f *Furb) ReqAPI() (interface{}, error) {
	var query string
	if f.Type == "manga" {
		query = apiURL + mangaAPI + q + f.Request
	} else if f.Type == "chapter" {
		query = apiURL + mangaAPI + chapterAPI + q + f.Request
	} else {
		fmt.Println(" [!] Invalid TYPE! Please do not change the type set in the code.")
		os.Exit(1) // stop the cli
	}

	// data resp
	var data interface{}

	resp, err := http.Get(query)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	unpk := json.Unmarshal(body, &data)
	if unpk != nil {
		return data, err
	}

	// return the json response
	return data, nil
}
