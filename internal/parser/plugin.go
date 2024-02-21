package parser

import (
	"context"
	"net/http"
	"net/url"
)

var plugins = make(map[string]Plugin)

type Plugin interface {
	Preprocess(context.Context, *url.URL) (*http.Request, error)

	Parse(context.Context, *http.Response) (*Feed, error)

	Postprocess(context.Context, *Feed) error
}

type InvalidURLError struct {
	URL url.URL
}

func RegisterPlugin(hostname string, plugin Plugin) {
	plugins[hostname] = plugin
}

func LoadPlugin(hostname string) Plugin {
	return plugins[hostname]
}
