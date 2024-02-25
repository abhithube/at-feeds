package parser

import (
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func HasMime(header http.Header, data []byte, options []string) bool {
	contentType := header.Get("Content-Type")
	if slices.ContainsFunc(options, func(s string) bool { return strings.Contains(contentType, s) }) {
		return true
	}

	mimeType := mimetype.Detect(data).String()
	return slices.ContainsFunc(options, func(s string) bool { return strings.Contains(mimeType, s) })
}

func GetNamedGroups(regex *regexp.Regexp, s string) map[string]string {
	match := regex.FindStringSubmatch(s)
	if match == nil {
		return nil
	}

	result := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
