package parser

import (
	"time"
)

type Feed struct {
	URL     string
	Link    string
	Title   string
	Entries []*Entry
}

type Entry struct {
	Link         string
	Title        string
	PublishedAt  time.Time
	Author       string
	Content      string
	ThumbnailURL string
}
