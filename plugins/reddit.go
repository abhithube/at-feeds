package plugins

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/abhithube/at-feeds/internal/parser"
)

type redditPlugin struct{}

func (p *redditPlugin) Preprocess(ctx context.Context, feedURL *url.URL) (*http.Request, error) {
	if !strings.Contains(feedURL.Path, ".rss") {
		feedURL.Path = feedURL.Path + ".rss"
	}

	return http.NewRequestWithContext(ctx, http.MethodGet, feedURL.String(), nil)
}

func (p *redditPlugin) Parse(ctx context.Context, resp *http.Response) (*parser.Feed, error) {
	return nil, nil
}

func (p *redditPlugin) Postprocess(ctx context.Context, feed *parser.Feed) error {
	return nil
}
