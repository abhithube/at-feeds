package parser

import (
	"regexp"
	"strings"
)

var imgRegex = regexp.MustCompile(`<img[^>]+src="([^"]+)"`)

type ThumbnailDescriptionProcessor struct{}

func NewThumbnailDescriptionProcessor() *ThumbnailDescriptionProcessor {
	return &ThumbnailDescriptionProcessor{}
}

func (p *ThumbnailDescriptionProcessor) Postprocess(feed *Feed) error {
	for _, entry := range feed.Entries {
		matches := imgRegex.FindStringSubmatch(entry.Content)
		if len(matches) == 0 {
			continue
		}

		match := matches[1]
		if !strings.HasPrefix(match, "https:") {
			match = "https:" + match
		}

		entry.ThumbnailURL = match
	}

	return nil
}
