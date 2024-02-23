package plugins

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/abhithube/at-feeds/internal/parser"
)

type youTubePlugin struct{}

var channelRegex = regexp.MustCompile(`/channel/(?P<channelID>UC[\w-_]+)`)

func (p *youTubePlugin) Preprocess(ctx context.Context, feedURL *url.URL) (*http.Request, error) {
	path := feedURL.Path
	if groups := parser.GetNamedGroups(channelRegex, path); groups != nil {
		feedURL.Path = "/feeds/videos.xml"
		feedURL.RawQuery = fmt.Sprintf("?channel_id=%s", groups["channelID"])
	}
	return http.NewRequestWithContext(ctx, http.MethodGet, feedURL.String(), nil)
}

func (p *youTubePlugin) Parse(ctx context.Context, resp *http.Response) (*parser.Feed, error) {
	return nil, nil
}

func (p *youTubePlugin) Postprocess(ctx context.Context, feed *parser.Feed) error {
	return nil
}
