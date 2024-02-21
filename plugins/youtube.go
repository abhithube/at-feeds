package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/abhithube/at-feeds/internal/parser"
)

type YouTubePlugin struct{}

var channelRegex = regexp.MustCompile(`/channel/(?P<channelID>UC[\w-_]+)`)

func (p *YouTubePlugin) Preprocess(ctx context.Context, feedURL *url.URL) (*http.Request, error) {
	path := feedURL.Path
	if groups := parser.GetNamedGroups(channelRegex, path); groups != nil {
		feedURL.Path = "/feeds/videos.xml"
		feedURL.RawQuery = fmt.Sprintf("?channel_id=%s", groups["channelID"])
	}
	return http.NewRequestWithContext(ctx, http.MethodGet, feedURL.String(), nil)
}

func (p *YouTubePlugin) Parse(ctx context.Context, resp *http.Response) (*parser.Feed, error) {
	return nil, nil
}

func (p *YouTubePlugin) Postprocess(ctx context.Context, feed *parser.Feed) error {
	return nil
}
