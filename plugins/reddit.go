package plugins

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/abhithube/at-feeds/internal/parser"
)

type RedditPlugin struct{}

func (p *RedditPlugin) Preprocess(ctx context.Context, feedURL *url.URL) (*http.Request, error) {
	if !strings.Contains(feedURL.Path, ".rss") {
		feedURL.Path = feedURL.Path + ".rss"
	}

	return http.NewRequestWithContext(ctx, http.MethodGet, feedURL.String(), nil)
}

func (p *RedditPlugin) Parse(ctx context.Context, resp *http.Response) (*parser.Feed, error) {
	return nil, nil
}

func (p *RedditPlugin) Postprocess(ctx context.Context, feed *parser.Feed) error {
	return nil
}
