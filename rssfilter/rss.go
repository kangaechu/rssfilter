package rssfilter

import (
	"fmt"
	"net/url"
	"regexp"
	"time"

	"github.com/gorilla/feeds"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/mmcdole/gofeed"
)

// RSS はRSS全体を示します
type RSS struct {
	Title   string      `json:"title,omitempty"`
	Entries *[]RSSEntry `json:"entries,omitempty"`
}

// Classify は未分類の記事を分類します
func (r RSS) Classify(classifier *BayesClassifier) error {
	for i, entry := range *r.Entries {
		if entry.Reputation != "" {
			continue
		}
		_, words, err := entry.GenerateLearnData()
		if err != nil {
			return err
		}
		cl, err := classifier.Classify(words)
		if err != nil {
			entry.Reputation = "ERROR"
		} else {
			entry.Reputation = cl
		}

		(*r.Entries)[i].Reputation = cl
		(*r.Entries)[i].LabeledBy = "auto"
	}
	return nil
}

// GenerateRss generates RSS XML
// URL: URL of the RSS
// publishPeriod: Period to publish
func (r RSS) GenerateRss(url string, publishAfter time.Time) (*string, error) {
	rssFeed := &feeds.Feed{
		Title:   r.Title,
		Link:    &feeds.Link{Href: url},
		Updated: time.Time{},
		Created: time.Time{},
	}
	for _, entry := range *r.Entries {
		if entry.Reputation != "Good" {
			continue
		}
		if entry.Published.Before(publishAfter) {
			continue
		}
		item := feeds.Item{
			Title:       entry.Title,
			Link:        &feeds.Link{Href: entry.Link},
			Description: entry.Description,
			Created:     entry.Published,
		}
		rssFeed.Add(&item)
	}
	rssXML, err := rssFeed.ToRss()
	if err != nil {
		return nil, err
	}
	return &rssXML, nil
}

// CreateRSSFromURL は指定されたURLからRSSを生成します。
func CreateRSSFromURL(URL string) (*RSS, error) {
	fp := gofeed.NewParser()
	rssfeeds, err := fp.ParseURL(URL)
	if err != nil {
		return nil, err
	}
	var entries []RSSEntry
	retrieved := time.Now()
	for _, item := range rssfeeds.Items {
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
	rss := RSS{Title: rssfeeds.Title, Entries: &entries}
	return &rss, nil
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
	LabeledBy   string    `json:"labeledBy,omitempty"`
}

// GenerateLearnData はRSSEntryから分類データを生成します。
func (e RSSEntry) GenerateLearnData() (string, *[]string, error) {

	var words []string
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return "", nil, err
	}
	posFilter := filter.NewPOSFilter([]filter.POS{{"名詞"}}...)
	// Title
	tokens := t.Tokenize(e.Title)
	posFilter.Keep(&tokens)

	ignoreCharacter := regexp.MustCompile(`^[\d./:,-]+$`)
	for _, token := range tokens {
		if len([]rune(token.Surface)) < 2 {
			continue
		}
		if ignoreCharacter.MatchString(token.Surface) {
			continue
		}
		words = append(words, token.Surface)
	}

	// Description
	tokens = t.Tokenize(e.Description)
	posFilter.Keep(&tokens)
	for _, token := range tokens {
		if len([]rune(token.Surface)) < 2 {
			continue
		}
		if ignoreCharacter.MatchString(token.Surface) {
			continue
		}
		words = append(words, token.Surface)
	}

	// URL(ドメイン名）
	parsedURL, err := url.Parse(e.Link)
	if err != nil {
		fmt.Println("error while parsing URL: ", e.Link, err)
	} else {
		words = append(words, "domain:"+parsedURL.Host)
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
