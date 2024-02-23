package parser

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
