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

	// 保存済みのRSSにないものは追加
	if oldRss == nil {
		oldRss = rss.Entries
	} else {
		for _, newEntry := range *rss.Entries {
			foundOldEntry := false
			for _, oldEntry := range *oldRss {
				if newEntry.Link == oldEntry.Link {
					foundOldEntry = true
					break
				}
			}
			if !foundOldEntry {
				*oldRss = append(*oldRss, newEntry)
			}
		}
	}

	newJson, err := json.Marshal(oldRss)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(j.FileName, newJson, 0644)
	if err != nil {
		return err
	}
	return nil
}
