package furb

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/TheBoringDude/furb-cli/utils"
)

// Download => chapter download
type Download struct {
	Furb
	Cwd        string
	Class      string
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

	if d.Class == "manga" {
		req, err := d.ReqAPI()
		utils.LogErr(err)

		resp := req.(map[string]interface{})

		imgDl.Images = resp["images"].([]interface{})
	} else if d.Class == "chapter" {
		imgDl.Images = d.DImg
	}

	// start downloading
	imgDl.imageDownload()
}

// download images
func (i *image) imageDownload() {
	var wg sync.WaitGroup

	fmt.Printf("\n\tDownloading => " + i.Title)

	for k, j := range i.Images {
		imgf := utils.StripFilename(k, j.(string))

		// start
		wg.Add(1)
		go worker(&wg, j.(string), i.ChapterCWD, imgf)
	}

	wg.Wait()

	fmt.Printf("\t...DONE\n")
}

func worker(wg *sync.WaitGroup, j string, cwd string, fname string) {
	defer wg.Done()

	resp, err := http.Get(j)
	utils.LogErr(err)

	defer resp.Body.Close()

	// create the file
	out, err := os.Create(cwd + "/" + fname)
	utils.LogErr(err)

	defer out.Close()

	// write resp body to file
	_, err = io.Copy(out, resp.Body)
}
