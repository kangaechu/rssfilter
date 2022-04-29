package rssfilter

import (
	"github.com/mmcdole/gofeed"
	"net/url"
	"time"
)

// RSS はRSS全体を示します
type RSS struct {
	Entries *[]RSSEntry
}

// RSSEntry はRSSの特定の記事を示します
type RSSEntry struct {
	Title       string
	Description string
	Link        url.URL
	Published   time.Time
	Categories  []string
}

// Import は指定されたURLからRSSを生成します。
func Import(URL string) (*RSS, error) {
	fp := gofeed.NewParser()
	feeds, err := fp.ParseURL(URL)
	if err != nil {
		return nil, err
	}
	var entries []RSSEntry
	for _, item := range feeds.Items {
		rssURL, _ := url.Parse(item.Link)
		published, _ := time.Parse(time.RFC3339, item.Published)
		var entry = RSSEntry{
			Title:       item.Title,
			Description: item.Description,
			Link:        *rssURL,
			Published:   published,
			Categories:  item.Categories,
		}
		entries = append(entries, entry)
	}
	rss := RSS{&entries}
	return &rss, nil
}
