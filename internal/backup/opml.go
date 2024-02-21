package backup

import (
	"encoding/xml"
)

type OPMLBackupManager struct{}

type OPMLDocument struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    OPMLHead `xml:"head"`
	Body    OPMLBody `xml:"body"`
}

type OPMLHead struct {
	Title string `xml:"title"`
}

type OPMLBody struct {
	Outline []OPMLOutline `xml:"outline"`
}

type OPMLOutline struct {
	Type    string `xml:"type,attr"`
	Text    string `xml:"text,attr"`
	HtmlURL string `xml:"htmlUrl,attr"`
	Title   string `xml:"title,attr"`
	XmlURL  string `xml:"xmlUrl,attr"`
}

func NewOPMLBackupManager() *OPMLBackupManager {
	return &OPMLBackupManager{}
}

func (bm *OPMLBackupManager) Import(data []byte) ([]BackupItem, error) {
	var doc OPMLDocument

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

func (bm *OPMLBackupManager) Export(items []BackupItem) ([]byte, error) {
	outlines := make([]OPMLOutline, len(items))

	for i := range items {
		item := items[i]
		outline := OPMLOutline{
			Type:    "rss",
			Text:    item.Title,
			HtmlURL: item.Link,
			Title:   item.Title,
			XmlURL:  item.URL,
		}
		outlines[i] = outline
	}

	doc := OPMLDocument{
		Version: "2.0",
		Head: OPMLHead{
			Title: "Feeds",
		},
		Body: OPMLBody{
			Outline: outlines,
		},
	}

	bytes, err := xml.Marshal(doc)

	return append([]byte(xml.Header), bytes...), err
}
