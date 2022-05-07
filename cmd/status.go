package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kangaechu/rssfilter/rssfilter"

	"github.com/spf13/cobra"
)

var statusModel string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of model",
	Long:  `status of model`,
	Run: func(cmd *cobra.Command, args []string) {
		// load model
		bayesClassifier, err := rssfilter.LoadBayesModel(statusModel)
		if err != nil {
			log.Fatal("failed to load model. err: ", err)
		}

		bayesStatus := bayesClassifier.Status()
		json, err := json.Marshal(bayesStatus)
		if err != nil {
			log.Fatal("failed to generate JSON of status err: ", err)
		}
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringVarP(&statusModel, "model", "m", "", "model file name")
	err := statusCmd.MarkFlagRequired("model")
	if err != nil {
		log.Fatal("specify model file name. err: ", err)
	}
}
