package rssfilter

import (
	"sort"

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
var classesStr = []string{"Good", "Bad"}

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

// BayesStatus stores the result of learned status
type BayesStatus struct {
	LearnedCount   int            `json:"learned_count,omitempty"`
	WordsByClasses []WordsByClass `json:"words_by_classes,omitempty"`
}

// WordsByClass stores class name and list of scores words.
type WordsByClass struct {
	ClassName  string      `json:"class_name,omitempty"`
	WordScores []WordScore `json:"word_scores,omitempty"`
}

// WordScore stores set of word and score
type WordScore struct {
	Word  string  `json:"word,omitempty"`
	Score float64 `json:"score,omitempty"`
}

// Status returns the result of learned status
func (b BayesClassifier) Status() *BayesStatus {
	var bayesStatus BayesStatus
	bayesStatus.LearnedCount = b.Classifier.Learned()
	for i, class := range classes {
		freq := b.Classifier.WordsByClass(class)
		words := sortWordsByScore(&freq)
		bayesStatus.WordsByClasses = append(bayesStatus.WordsByClasses, WordsByClass{
			ClassName: classesStr[i], WordScores: (*words)[0:30]})
	}
	return &bayesStatus
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

// sortWordsByScore returns WordScore that is sorted by score desc.
func sortWordsByScore(freqMap *map[string]float64) *[]WordScore {
	words := make([]WordScore, 0, len(*freqMap))
	for k, v := range *freqMap {
		words = append(words, WordScore{Word: k, Score: v})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].Score > words[j].Score
	})
	return &words
}
