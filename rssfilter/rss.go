package rssfilter

import (
	"github.com/mmcdole/gofeed"
	"time"
)

// RSS はRSS全体を示します
type RSS struct {
	Entries *[]RSSEntry
}

// RSSEntry はRSSの特定の記事を示します
type RSSEntry struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Link        string    `json:"link"`
	Published   time.Time `json:"published"`
	Categories  []string  `json:"categories,omitempty"`
	Retrieved   time.Time `json:"retrieved,omitempty"`
	Reputation  string    `json:"reputation,omitempty"`
}

// Import は指定されたURLからRSSを生成します。
func Import(URL string) (*RSS, error) {
	fp := gofeed.NewParser()
	feeds, err := fp.ParseURL(URL)
	if err != nil {
		return nil, err
	}
	var entries []RSSEntry
	retrieved := time.Now()
	for _, item := range feeds.Items {
		published, _ := time.Parse(time.RFC3339, item.Published)
		var entry = RSSEntry{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Published:   published,
			Categories:  item.Categories,
			Retrieved:   retrieved,
		}
		entries = append(entries, entry)
	}
	rss := RSS{&entries}
	return &rss, nil
}
