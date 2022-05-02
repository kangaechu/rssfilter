package cmd

import (
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
)

var learnRSSJSON string
var learnModel string

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "create / update classifier",
	Long:  `create / update classifier`,
	Run: func(cmd *cobra.Command, args []string) {
		storageJSON := rssfilter.StorageJSON{FileName: learnRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON. json:", learnRSSJSON, ", err: ", err)
		}

		model, err := rssfilter.GenerateBayesModel(*rss)
		if err != nil {
			log.Fatal("failed to generate model. err: ", err)
		}

		err = model.Store(learnModel)
		if err != nil {
			log.Fatal("failed to store model. model: ", learnModel, ", err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
	learnCmd.Flags().StringVarP(&learnRSSJSON, "feed", "f", "", "feed JSON file name")
	err := learnCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}

	learnCmd.Flags().StringVarP(&learnModel, "model", "m", "", "model file name")
	err = learnCmd.MarkFlagRequired("model")
	if err != nil {
		log.Fatal("specify model file name. err: ", err)
	}
}
