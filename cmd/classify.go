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
	Run: func(cmd *cobra.Command, args []string) {
		// JSONを読み込む
		storageJSON := rssfilter.StorageJSON{FileName: classifyRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON ", classifyRSSJSON, err)
		}

		// モデルを読み込む
		bayesClassifier, err := rssfilter.LoadBayesModel(classifyModel)
		if err != nil {
			log.Fatal("failed to load model", err)
		}

		// 分類
		err = rss.Classify(bayesClassifier)
		if err != nil {
			log.Fatal("failed to store model", learnModel, err)
		}

		// JSONを保存
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to store JSON", classifyRSSJSON, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(classifyCmd)
	rootCmd.AddCommand(learnCmd)
	classifyCmd.Flags().StringVarP(&classifyRSSJSON, "feed", "f", "", "feed JSON file name")
	err := classifyCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name")
	}

	classifyCmd.Flags().StringVarP(&classifyModel, "model", "m", "", "model file name")
	err = classifyCmd.MarkFlagRequired("model")
	if err != nil {
		log.Fatal("specify model file name")
	}
}
