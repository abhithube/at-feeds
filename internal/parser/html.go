package parser

import (
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

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
