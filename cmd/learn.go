package cmd

import (
	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
	"log"
)

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "create / update classifier",
	Long:  `create / update classifier`,
	Run: func(cmd *cobra.Command, args []string) {
		storageJSON := rssfilter.StorageJSON{FileName: "hoge.json"}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON ", err)
		}

		model, err := rssfilter.GenerateBayesModel(*rss)
		if err != nil {
			log.Fatal("failed to generate model", err)
		}

		err = model.Store("model")
		if err != nil {
			log.Fatal("failed to store model", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// learnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// learnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
