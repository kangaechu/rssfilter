package rssfilter

import (
	"fmt"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
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

// GenerateLearnData はRSSEntryから分類データを生成します。
func (e RSSEntry) GenerateLearnData() (string, *[]string, error) {

	var words []string
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return "", nil, err
	}
	posFilter := filter.NewPOSFilter([]filter.POS{{"名詞"}, {"形容詞"}}...)
	// Title
	tokens := t.Tokenize(e.Title)
	posFilter.Keep(&tokens)
	for _, token := range tokens {
		words = append(words, token.Surface)
	}

	// Description
	tokens = t.Tokenize(e.Description)
	posFilter.Keep(&tokens)
	for _, token := range tokens {
		words = append(words, token.Surface)
	}

	// URL(ドメイン名）
	parsedURL, err := url.Parse(e.Link)
	if err != nil {
		fmt.Println("error while parsing URL: ", e.Link, err)
	} else {
		words = append(words, parsedURL.Host)
	}

	// タグ
	for _, category := range e.Categories {
		words = append(words, "category:"+category)
	}

	uniqueWords := SliceUnique(words)
	return e.Reputation, &uniqueWords, nil
}

// SliceUnique はリストの重複を取り除きます
func SliceUnique(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
