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

// chapterCmd represents the chapter command
var chapterCmd = &cobra.Command{
	Use:   "chapter",
	Short: "Download a specific chapter of a manga.",
	Long: `
Download a specific chapter of a manga from a website.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// initialize new furb downloader
		session := furb.Furb{
			Request: qManga,
			Type:    "chapter",
		}

		// validate request
		// it will exit on its own, upon error
		session.InitConf()

		// request the manga api
		rs, err := session.ReqAPI()
		utils.LogErr(err)

		resp := rs.(map[string]interface{})

		fmt.Println(resp["title"])

		// get the current working dir
		cwd, err := os.Getwd()
		utils.LogErr(err)

		// initiate download
		download := furb.Download{
			Furb:  session,
			Cwd:   cwd,
			Title: resp["title"].(string),
			DImg:  resp["images"].([]interface{}),
		}

		// download
		go download.DownloadChapter()
	},
}

func init() {
	downloadCmd.AddCommand(chapterCmd)
	chapterCmd.Flags().StringVarP(&qManga, "site-chapter", "s", "", "Manga, manhuwa, manhua - chapter website link.")
	chapterCmd.MarkFlagRequired("site-chapter")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chapterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chapterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
