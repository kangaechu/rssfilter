package rssfilter

// BayesModel はナイーブベイズのモデルです
type BayesModel struct {
	Model map[string]*[]string
}

// Store はナイーブベイズのモデルを保存します
func (b BayesModel) Store() {

}

// GenerateBayesModel はRSSからナイーブベイズのモデルを生成します。
func GenerateBayesModel(r RSS) (*BayesModel, error) {
	var bayesModel BayesModel
	//
	//for _, entry := range *r.Entries {
	//  className, words, err := entry.GenerateModel()
	//  if err != nil {
	//    return nil, err
	//  }
	//  *bayesModel.Model[className] = append(*bayesModel.Model[className], *words...)
	//}
	return &bayesModel, nil
}
