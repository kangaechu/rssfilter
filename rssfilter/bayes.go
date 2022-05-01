package rssfilter

import "github.com/navossoc/bayesian"

// BayesClassifier はナイーブベイズのモデルです
type BayesClassifier struct {
	Classifier *bayesian.Classifier
}

// 分類するクラス一覧
const (
	Good bayesian.Class = "Good"
	Bad  bayesian.Class = "Bad"
)

// Store はナイーブベイズのモデルを保存します
func (b BayesClassifier) Store(filename string) error {
	err := b.Classifier.WriteToFile(filename)
	if err != nil {
		return err
	}
	return nil
}

// GenerateBayesModel はRSSからナイーブベイズのモデルを生成します。
func GenerateBayesModel(r RSS) (*BayesClassifier, error) {
	var bayesClassifier BayesClassifier
	bayesClassifier.Classifier = bayesian.NewClassifier(Good, Bad)

	for _, entry := range *r.Entries {
		if entry.Reputation == "" {
			continue
		}
		className, words, err := entry.GenerateLearnData()
		if err != nil {
			return nil, err
		}
		var cl bayesian.Class
		if className == "Good" {
			cl = Good
		} else {
			cl = Bad
		}
		bayesClassifier.Classifier.Learn(*words, cl)
	}
	return &bayesClassifier, nil
}
