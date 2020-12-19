package furb

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/TheBoringDude/furb-cli/utils"
)

// Download => chapter download
type Download struct {
	Furb
	Cwd        string
	Title      string
	ChapterURL string
	DImg       []interface{} // this is only used by DownloadChapter
}

// image fields
type image struct {
	ChapterCWD string
	Title      string
	Images     []interface{}
}

// DownloadChapter downloads the images in the set chapter of the manga.
func (d *Download) DownloadChapter() {
	chapterDir := d.Cwd + "/" + d.Title

	// make the chapter dir
	utils.MakeDir(chapterDir)

	// initialiaze image grabber
	imgDl := image{
		ChapterCWD: chapterDir,
		Title:      d.Title,
	}

	if d.ChapterURL != "" && len(d.DImg) > 0 {
		req, err := d.ReqAPI()
		utils.LogErr(err)

		resp := req.(map[string]interface{})

		imgDl.Images = resp["images"].([]interface{})
	} else {
		imgDl.Images = d.DImg
	}

	// start downloading
	go imgDl.imageDownload()
}

// download images
func (i *image) imageDownload() {
	fmt.Println("Downloading => " + i.Title)

	for k, j := range i.Images {
		imgf := utils.StripFilename(k, j.(string))

		resp, err := http.Get(j.(string))
		utils.LogErr(err)

		defer resp.Body.Close()

		// create the file
		out, err := os.Create(i.ChapterCWD + "/" + imgf)
		utils.LogErr(err)

		defer out.Close()

		// write resp body to file
		_, err = io.Copy(out, resp.Body)
	}
}
