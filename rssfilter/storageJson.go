package rssfilter

import (
	"encoding/json"
	"io/ioutil"
)

// StorageJSON はStorageを実装し、JSONによりRSSを保存します。
type StorageJSON struct {
	FileName string
}

func (j StorageJSON) Store(rss *RSS) error {
	rssJson, err := json.Marshal(rss.Entries)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(j.FileName, rssJson, 0644)
	if err != nil {
		return err
	}
	return nil
}
