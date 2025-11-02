package rssfilter

// Storage is an interface for storing RSS feed data.
type Storage interface {
	Store(rss *RSS)
}
