package rssfilter

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// StorageJSON はStorageを実装し、JSONによりRSSを保存します。
type StorageJSON struct {
	FileName string
}

// StoreUnique はRSSと既に保存済みのファイルを比較し、新たに追加されたものを追記します
func (j StorageJSON) StoreUnique(rss *RSS) error {

	var oldRss *[]RSSEntry
	if f, err := os.Stat(j.FileName); os.IsNotExist(err) || f.IsDir() {
	} else {
		// ファイルが存在するときは読み込む
		oldRssText, err := ioutil.ReadFile(j.FileName)
		if err != nil {
			return err
		}
		err = json.Unmarshal(oldRssText, &oldRss)
		if err != nil {
			return err
		}
	}

	mergedRss := merge(oldRss, rss.Entries)
	newJson, err := json.Marshal(mergedRss)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(j.FileName, newJson, 0644)
	if err != nil {
		return err
	}
	return nil
}

// merge はLinkをキーにして2つのリストをマージします
func merge(rss1, rss2 *[]RSSEntry) *[]RSSEntry {
	if rss1 == nil {
		return rss2
	}
	if rss2 == nil {
		return rss1
	}
	mergedRss := rss1
	for _, entry2 := range *rss2 {
		found := false
		for _, entry1 := range *rss1 {
			if entry2.Link == entry1.Link {
				found = true
				break
			}
		}
		if !found {
			*mergedRss = append(*mergedRss, entry2)
		}
	}
	return mergedRss
}
