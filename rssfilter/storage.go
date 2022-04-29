package rssfilter

type Storage interface {
	Store(rss *RSS)
}
