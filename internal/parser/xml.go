package parser

import (
	"encoding/xml"
	"time"
)

type RSSFeed struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title string    `xml:"title"`
	Link  string    `xml:"link"`
	Items []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	PubDate     string        `xml:"pubDate"`
	Description string        `xml:"description"`
	Author      string        `xml:"author"`
	Enclosure   *RSSEnclosure `xml:"enclosure"`
	Thumbnail   string
	DC
}

type RSSEnclosure struct {
	URL    string `xml:"url,attr"`
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type AtomFeed struct {
	Feed    AtomName    `xml:"feed"`
	Title   string      `xml:"title"`
	Links   []AtomLink  `xml:"link"`
	Entries []AtomEntry `xml:"entry"`
}

type AtomName struct {
	Feed  xml.Name `xml:"feed"`
	Media string   `xml:"xmlns:media,attr"`
}

type AtomEntry struct {
	Title     string       `xml:"title"`
	Links     []AtomLink   `xml:"link"`
	Authors   []AtomAuthor `xml:"author"`
	Published time.Time    `xml:"published"`
	Summary   string       `xml:"summary"`
	Media
}

type AtomLink struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type AtomAuthor struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

type Media struct {
	Group     *MediaGroup     `xml:"group"`
	Thumbnail *MediaThumbnail `xml:"thumbnail"`
}

type MediaGroup struct {
	Title       string          `xml:"title"`
	Description string          `xml:"description"`
	Thumbnail   *MediaThumbnail `xml:"thumbnail"`
}

type MediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type DC struct {
	Creator string `xml:"creator"`
}
