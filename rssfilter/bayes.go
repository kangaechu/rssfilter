package rssfilter

import (
	"github.com/navossoc/bayesian"
)

// BayesClassifier はナイーブベイズのモデルです
type BayesClassifier struct {
	Classifier *bayesian.Classifier
}

// 分類するクラス一覧
const (
	Good bayesian.Class = "Good"
	Bad  bayesian.Class = "Bad"
)

var classes = []bayesian.Class{Good, Bad}

// Store はナイーブベイズのモデルを保存します
func (b BayesClassifier) Store(filename string) error {
	err := b.Classifier.WriteToFile(filename)
	if err != nil {
		return err
	}
	return nil
}

func (b BayesClassifier) Classify(words *[]string) (string, error) {
	_, inx, _, err := b.Classifier.SafeProbScores(*words)
	if err != nil {
		return "", err
	}
	var cl string
	if inx == 0 {
		cl = "Good"
	} else if inx == 1 {
		cl = "Bad"
	} else {
		cl = "Undefined"
	}
	return cl, nil
}

// GenerateBayesModel はRSSからナイーブベイズのモデルを生成します。
func GenerateBayesModel(r RSS) (*BayesClassifier, error) {
	var bayesClassifier BayesClassifier
	bayesClassifier.Classifier = bayesian.NewClassifier(classes...)

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

// LoadBayesModel はファイルからモデルをロードします
func LoadBayesModel(filename string) (*BayesClassifier, error) {
	c, err := bayesian.NewClassifierFromFile(filename)
	if err != nil {
		return nil, err
	}
	bc := BayesClassifier{Classifier: c}

	return &bc, nil
}
