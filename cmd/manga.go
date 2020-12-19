/*Package cmd ...
Copyright © 2020 TheBoringDude <iamcoderx@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/TheBoringDude/furb-cli/furb"
	"github.com/TheBoringDude/furb-cli/utils"
	"github.com/spf13/cobra"
)

// qManga -> the request manga
var qManga string

// mangaCmd represents the manga command
var mangaCmd = &cobra.Command{
	Use:   "manga -s <manga-url-website>",
	Short: "Download a full manga.",
	Long: `
Download all chapters of a manga, manhuwa, manhua from a specific website.

It overrides any existing folder and chapter folders depend on the name of the
chapter from the website.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// initialize new furb downloader
		session := furb.Furb{
			Request: qManga,
			Type:    "manga",
		}

		// validate request
		// it will exit on its own, upon error
		session.InitConf()

		// request the manga api
		rs, err := session.ReqAPI()
		utils.LogErr(err)

		resp := rs.(map[string]interface{})

		fmt.Println(resp["title"]) // print the manga title

		// get the current working dir
		cwd, err := os.Getwd()
		utils.LogErr(err)

		mDir := cwd + "/" + resp["title"].(string)

		// create the dir
		utils.MakeDir(mDir)

		// reverse the slice since it starts from the latest chapter to the earliest
		chapters := utils.ReverseSlice(resp["chapters"].([]interface{}))

		// extract chapters
		for _, ch := range chapters {
			chapter := ch.(map[string]interface{})

			// init new download
			download := furb.Download{
				Furb: furb.Furb{
					Request: chapter["chapter_url"].(string),
					Type:    "chapter",
				},
				Cwd:        mDir,
				Title:      chapter["chapter_name"].(string),
				ChapterURL: chapter["chapter_url"].(string),
			}

			// {session, mDir, chapter["chapter_name"].(string), chapter["chapter_url"].(string)}
			// download each
			go download.DownloadChapter()
		}
	},
}

func init() {
	downloadCmd.AddCommand(mangaCmd)
	mangaCmd.Flags().StringVarP(&qManga, "site", "s", "", "Manga, manhuwa, manhua website link.")
	mangaCmd.MarkFlagRequired("site")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mangaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mangaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
