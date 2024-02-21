package plugins

import "github.com/abhithube/at-feeds/internal/parser"

func Register() {
	parser.RegisterPlugin("www.reddit.com", &RedditPlugin{})
	parser.RegisterPlugin("www.youtube.com", &YouTubePlugin{})
}
