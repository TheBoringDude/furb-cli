package furb

import (
	"os"
	"log"
	"fmt"
	"strings"
	"net/http"
	"io"
	"github.com/TheBoringDude/furb-cli/utils"
)

// DownloadChapter downloads the images in a chapter.
func DownloadChapter(cwd, title, chapterURL string){
	req, err := utils.Request(chapterURL, "chapter")
	if err != nil {
		log.Fatalln(err) // log the error
	}

	resp := req.(map[string]interface{})

	chapterDir := cwd + "/" + title

	// make the chapter dir
	err = os.Mkdir(chapterDir, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	go imageDownload(chapterDir, title,  resp["images"].([]interface{}))
}

// download images
func imageDownload(chapterCwd string, title string, images []interface{}) {
	fmt.Println("Downloading => " + title)

	for k, i := range images{
		imgf := stripFilename(k, i.(string))

		resp, err := http.Get(i.(string))
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		// create the file
		out, err := os.Create(chapterCwd + "/" + imgf)
		if err != nil {
			log.Fatalln(err)
		}

		defer out.Close()

		// write resp body to file
		_, err = io.Copy(out, resp.Body)
	}
}

// strip filename
func stripFilename(count int, imgURL string) string{
	// get the filename in the url
	strs := strings.Split(imgURL, "/")

	// rename it
	ren := strings.Split(strs[len(strs)-1], ".")

	// return new filename
	return fmt.Sprintf("%03d." + ren[1], count)
}