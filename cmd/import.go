package cmd

import (
	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
	"log"
)

var RSSURL string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Retrieve RSS and store",
	Long:  `Retrieve RSS and store`,
	Run: func(cmd *cobra.Command, args []string) {

		// URLをもとにRSSを生成
		rss, err := rssfilter.Import(RSSURL)
		if err != nil {
			log.Fatal("failed to Parse Rss", RSSURL)
		}
		// 保存
		var storageJSON = rssfilter.StorageJSON{FileName: "hoge.json"}
		err = storageJSON.Store(rss)
		if err != nil {
			log.Fatal("failed to save json", RSSURL)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&RSSURL, "url", "u", "", "URL for retrieve URL")
	err := importCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatal("specify url")
	}
}
