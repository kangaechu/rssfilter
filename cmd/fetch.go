package cmd

import (
	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
	"log"
)

var fetchRSSURL string
var fetchRSSJSON string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieve RSS and store",
	Long:  `Retrieve RSS and store`,
	Run: func(cmd *cobra.Command, args []string) {

		// URLをもとにRSSを生成
		rss, err := rssfilter.CreateRSSFromURL(fetchRSSURL)
		if err != nil {
			log.Fatal("failed to parse rss. URL: ", fetchRSSURL, ", err: ", err)
		}
		// 保存
		var storageJSON = rssfilter.StorageJSON{FileName: fetchRSSJSON}
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to save json. json: ", fetchRSSJSON, ", err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&fetchRSSURL, "url", "u", "", "URL for retrieve URL")
	err := fetchCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatal("specify url. err: ", err)
	}

	fetchCmd.Flags().StringVarP(&fetchRSSJSON, "feed", "f", "", "feed JSON file name")
	err = fetchCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}
}
