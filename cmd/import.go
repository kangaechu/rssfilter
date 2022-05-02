package cmd

import (
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
)

var importRSSURL string
var importRSSJSON string

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Retrieve RSS and store",
	Long:  `Retrieve RSS and store`,
	Run: func(cmd *cobra.Command, args []string) {

		// URLをもとにRSSを生成
		rss, err := rssfilter.Import(importRSSURL)
		if err != nil {
			log.Fatal("failed to parse rss ", importRSSURL, err)
		}
		// 保存
		var storageJSON = rssfilter.StorageJSON{FileName: importRSSJSON}
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to save json ", importRSSJSON, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&importRSSURL, "url", "u", "", "URL for retrieve URL")
	err := importCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatal("specify url")
	}

	importCmd.Flags().StringVarP(&importRSSJSON, "feed", "f", "", "feed JSON file name")
	err = importCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name")
	}
}
