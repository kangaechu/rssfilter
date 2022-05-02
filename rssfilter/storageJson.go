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

// Load は保存済みのファイルをロードします
func (j StorageJSON) Load() (*RSS, error) {
	var rss RSS
	if f, err := os.Stat(j.FileName); os.IsNotExist(err) || f.IsDir() {
	} else {
		// ファイルが存在するときは読み込む
		oldRssText, err := ioutil.ReadFile(j.FileName)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(oldRssText, &rss.Entries)
		if err != nil {
			return nil, err
		}
	}
	return &rss, nil
}

// StoreUnique はRSSと既に保存済みのファイルを比較し、新たに追加されたものを追記します
func (j StorageJSON) StoreUnique(rss *RSS) error {
	oldRss, err := j.Load()
	if err != nil {
		return err
	}

	mergedRss := merge(oldRss.Entries, rss.Entries)
	newJson, err := json.Marshal(mergedRss)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(j.FileName, newJson, 0600)
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
		for i1, entry1 := range *rss1 {
			if entry2.Link == entry1.Link {
				found = true
				(*mergedRss)[i1] = entry2
				break
			}
		}
		if !found {
			*mergedRss = append(*mergedRss, entry2)
		}
	}
	return mergedRss
}
