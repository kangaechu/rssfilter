package cmd

import (
	"io/ioutil"
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"

	"github.com/spf13/cobra"
)

var exportRSSJSON string
var exportRSSXML string
var exportRSSURL string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export to RSS file",
	Long:  `export to RSS file`,
	Run: func(_ *cobra.Command, _ []string) {
		// Open feed JSON
		storageJSON := rssfilter.StorageJSON{FileName: exportRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON. json:", exportRSSJSON, ", err: ", err)
		}

		// generate RSS
		rssXML, err := rss.GenerateRss(exportRSSURL)
		if err != nil {
			log.Fatal("failed to generate RSS. err: ", err)
		}

		// store RSS
		err = ioutil.WriteFile(exportRSSXML, []byte(*rssXML), 0600)
		if err != nil {
			log.Fatal("failed to save RSS. err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringVarP(&exportRSSJSON, "feed", "f", "", "feed JSON file name")
	err := exportCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}

	exportCmd.Flags().StringVarP(&exportRSSXML, "rss", "r", "", "RSS xml file name")
	err = exportCmd.MarkFlagRequired("rss")
	if err != nil {
		log.Fatal("specify RSS xml file name. err: ", err)
	}

	exportCmd.Flags().StringVarP(&exportRSSURL, "url", "u", "https://example.com/rss.xml", "RSS URL Link")
}
