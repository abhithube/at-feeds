package parser

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

var atomMimeTypes = []string{"application/atom+xml"}
var rssMimeTypes = []string{"application/rss+xml"}
var htmlMimeTypes = []string{"text/html", "application/xhtml+xml"}

func HasMime(header http.Header, data []byte, options []string) bool {
	contentType := header.Get("Content-Type")
	if slices.ContainsFunc(options, func(s string) bool { return strings.Contains(contentType, s) }) {
		return true
	}

	mimeType := mimetype.Detect(data).String()
	return slices.ContainsFunc(options, func(s string) bool { return strings.Contains(mimeType, s) })
}

func IsHTMLDocument(header http.Header, data []byte) bool {
	return HasMime(header, data, htmlMimeTypes)
}

func IsAtomFeed(header http.Header, data []byte) bool {
	return HasMime(header, data, atomMimeTypes)
}

func IsRSSFeed(header http.Header, data []byte) bool {
	return HasMime(header, data, rssMimeTypes)
}
