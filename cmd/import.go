package cmd

import (
	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
	"log"
)

var RSSURL string
var RSSJSON string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Retrieve RSS and store",
	Long:  `Retrieve RSS and store`,
	Run: func(cmd *cobra.Command, args []string) {

		// URLをもとにRSSを生成
		rss, err := rssfilter.Import(RSSURL)
		if err != nil {
			log.Fatal("failed to parse rss ", RSSURL, err)
		}
		// 保存
		var storageJSON = rssfilter.StorageJSON{FileName: RSSJSON}
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to save json ", RSSJSON, err)
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

	importCmd.Flags().StringVarP(&RSSJSON, "feed", "f", "", "feed JSON file name")
	err = importCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name")
	}
}
