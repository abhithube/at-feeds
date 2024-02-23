package backup

import (
	"encoding/xml"
)

type OPMLManager struct{}

type opmlDocument struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    opmlHead `xml:"head"`
	Body    opmlBody `xml:"body"`
}

type opmlHead struct {
	Title string `xml:"title"`
}

type opmlBody struct {
	Outline []opmlOutline `xml:"outline"`
}

type opmlOutline struct {
	Type    string `xml:"type,attr"`
	Text    string `xml:"text,attr"`
	HtmlURL string `xml:"htmlUrl,attr"`
	Title   string `xml:"title,attr"`
	XmlURL  string `xml:"xmlUrl,attr"`
}

func NewOPMLManager() *OPMLManager {
	return &OPMLManager{}
}

func (m *OPMLManager) Import(data []byte) ([]BackupItem, error) {
	var doc opmlDocument

	if err := xml.Unmarshal(data, &doc); err != nil {
		return nil, err
	}

	outlines := doc.Body.Outline

	items := make([]BackupItem, len(outlines))

	for i := range outlines {
		outline := outlines[i]
		item := BackupItem{
			URL:   outline.XmlURL,
			Link:  outline.HtmlURL,
			Title: outline.Title,
		}
		items[i] = item
	}

	return items, nil
}

func (m *OPMLManager) Export(items []BackupItem) ([]byte, error) {
	outlines := make([]opmlOutline, len(items))

	for i := range items {
		item := items[i]
		outline := opmlOutline{
			Type:    "rss",
			Text:    item.Title,
			HtmlURL: item.Link,
			Title:   item.Title,
			XmlURL:  item.URL,
		}
		outlines[i] = outline
	}

	doc := opmlDocument{
		Version: "2.0",
		Head: opmlHead{
			Title: "Feeds",
		},
		Body: opmlBody{
			Outline: outlines,
		},
	}

	bytes, err := xml.Marshal(doc)

	return append([]byte(xml.Header), bytes...), err
}
