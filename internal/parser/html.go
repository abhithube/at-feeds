package parser

import (
	"net/http"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func IsHTMLDocument(header http.Header, data []byte) bool {
	return HasMime(header, data, []string{"text/html", "application/xhtml+xml"})
}

func Query(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return nil
	}

	return cascadia.Query(n, sel)
}

func QueryAll(n *html.Node, query string) []*html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return nil
	}

	return cascadia.QueryAll(n, sel)
}

func Attr(n *html.Node, attrName string) string {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}

	return ""
}
