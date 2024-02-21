package parser

import (
	"regexp"
	"strings"
)

type ThumbnailDescriptionProcessor struct{}

func NewThumbnailDescriptionProcessor() *ThumbnailDescriptionProcessor {
	return &ThumbnailDescriptionProcessor{}
}

func (p *ThumbnailDescriptionProcessor) Postprocess(feed *Feed) error {
	re := regexp.MustCompile("<img[^>]+src=\"([^\"]+)\"")

	for i := range feed.Entries {
		matches := re.FindStringSubmatch(feed.Entries[i].Content)
		if len(matches) == 0 {
			continue
		}

		match := matches[1]
		if !strings.HasPrefix(match, "https:") {
			match = "https:" + match
		}

		feed.Entries[i].ThumbnailURL = match
	}

	return nil
}
