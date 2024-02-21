package parser

import (
	"encoding/xml"
	"strings"
	"time"
)

type RSSParser struct{}

func NewRSSParser() *RSSParser {
	return &RSSParser{}
}

func (p *RSSParser) Parse(data []byte) (*Feed, error) {
	var feed *RSSFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	entries := make([]*Entry, len(feed.Channel.Items))
	for i := range feed.Channel.Items {
		item := feed.Channel.Items[i]

		publishedAt, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			publishedAt, _ = time.Parse("Mon 02 Jan 2006 15:04:05 -0700", item.PubDate)
		}

		entry := Entry{
			Link:        item.Link,
			Title:       strings.TrimSpace(item.Title),
			PublishedAt: publishedAt,
			Author:      item.Author,
			Content:     strings.TrimSpace(item.Description),
		}
		if item.Creator != "" {
			entry.Author = item.Creator
		}

		if item.Enclosure != nil && strings.HasPrefix(item.Enclosure.Type, "image/") {
			entry.ThumbnailURL = item.Enclosure.URL
		}

		entries[i] = &entry
	}

	parsed := &Feed{
		Link:    feed.Channel.Link,
		Title:   strings.TrimSpace(feed.Channel.Title),
		Entries: entries,
	}

	return parsed, nil
}
