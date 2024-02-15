package cmd

import (
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
)

var trainRSSJSON string
var trainModel string

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "train classifier",
	Long:  `train classifier`,
	Run: func(_ *cobra.Command, _ []string) {
		storageJSON := rssfilter.StorageJSON{FileName: trainRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON. json:", trainRSSJSON, ", err: ", err)
		}

		model, err := rssfilter.GenerateBayesModel(*rss)
		if err != nil {
			log.Fatal("failed to generate model. err: ", err)
		}

		err = model.Store(trainModel)
		if err != nil {
			log.Fatal("failed to store model. model: ", trainModel, ", err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
	trainCmd.Flags().StringVarP(&trainRSSJSON, "feed", "f", "", "feed JSON file name")
	err := trainCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}

	trainCmd.Flags().StringVarP(&trainModel, "model", "m", "", "model file name")
	err = trainCmd.MarkFlagRequired("model")
	if err != nil {
		log.Fatal("specify model file name. err: ", err)
	}
}
