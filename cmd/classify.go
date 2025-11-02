// Package cmd provides command-line interface functionality for rssfilter.
package cmd

import (
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"

	"github.com/spf13/cobra"
)

var classifyRSSJSON string
var classifyModel string

// classifyCmd represents the classify command
var classifyCmd = &cobra.Command{
	Use:   "classify",
	Short: "classify unclassified items",
	Long:  `classify unclassified items`,
	Run: func(_ *cobra.Command, _ []string) {
		// JSONを読み込む
		storageJSON := rssfilter.StorageJSON{FileName: classifyRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON. json:", classifyRSSJSON, ", err: ", err)
		}

		// モデルを読み込む
		bayesClassifier, err := rssfilter.LoadBayesModel(classifyModel)
		if err != nil {
			log.Fatal("failed to load model. err: ", err)
		}

		// 分類
		err = rss.Classify(bayesClassifier)
		if err != nil {
			log.Fatal("failed to classify. err: ", err)
		}

		// JSONを保存
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to store JSON. json: ", classifyRSSJSON, ", err: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(classifyCmd)
	classifyCmd.Flags().StringVarP(&classifyRSSJSON, "feed", "f", "", "feed JSON file name")
	err := classifyCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}

	classifyCmd.Flags().StringVarP(&classifyModel, "model", "m", "", "model file name")
	err = classifyCmd.MarkFlagRequired("model")
	if err != nil {
		log.Fatal("specify model file name. err: ", err)
	}
}
