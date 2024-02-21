package parser

import (
	"encoding/xml"
	"strings"
)

type AtomParser struct{}

func NewAtomParser() *AtomParser {
	return &AtomParser{}
}

func (p *AtomParser) Parse(data []byte) (*Feed, error) {
	var feed *AtomFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	entries := make([]*Entry, len(feed.Entries))
	for i, item := range feed.Entries {
		link := getLink(item.Links)
		if link == "" {
			continue
		}

		entry := &Entry{
			Link:        getLink(item.Links),
			Title:       strings.TrimSpace(item.Title),
			PublishedAt: item.Published,
			Content:     strings.TrimSpace(item.Summary),
		}
		if len(item.Authors) > 0 {
			entry.Author = item.Authors[0].Name
		}
		if item.Thumbnail != nil {
			entry.ThumbnailURL = item.Thumbnail.URL
		}

		if group := item.Media.Group; group != nil {
			if group.Title != "" {
				entry.Title = strings.TrimSpace(group.Title)
			}
			if group.Description != "" {
				entry.Content = strings.TrimSpace(group.Description)
			}
			if group.Thumbnail != nil {
				entry.ThumbnailURL = group.Thumbnail.URL
			}
		}

		entries[i] = entry
	}

	parsed := &Feed{
		Link:    getLink(feed.Links),
		Title:   strings.TrimSpace(feed.Title),
		Entries: entries,
	}

	return parsed, nil
}

func getLink(links []AtomLink) string {
	if len(links) == 1 {
		return links[0].Href
	}

	for _, item := range links {
		if item.Rel == "alternate" {
			return item.Href
		}
	}

	return ""
}
