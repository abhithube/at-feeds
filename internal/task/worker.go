package task

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/parser"
)

type Worker struct {
	manager *Manager
}

func NewWorker(manager *Manager) *Worker {
	return &Worker{manager: manager}
}

const limit = 5

func (w *Worker) RunAll(ctx context.Context, feedURLs []string) error {
	var wg sync.WaitGroup
	queue := make(chan struct{}, limit)

	for _, url := range feedURLs {
		queue <- struct{}{}
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			log.Printf("Starting %s...\n", url)

			if _, err := w.Run(ctx, url); err == nil {
				log.Printf("Finished %s!\n", url)
			} else {
				log.Printf("%s: %s", url, err)
			}

			<-queue
		}(url)
	}

	wg.Wait()
	close(queue)

	return nil
}

func (w *Worker) Run(ctx context.Context, feedURL string) (*database.Feed, error) {
	parsedURL, err := url.Parse(feedURL)
	if err != nil {
		return nil, err
	}

	plugin := parser.LoadPlugin(parsedURL.Hostname())

	var req *http.Request
	if plugin != nil {
		req, err = plugin.Preprocess(ctx, parsedURL)
	}
	if req == nil {
		req, err = w.manager.Preprocess(ctx, parsedURL)
	}
	if err != nil {
		return nil, err
	}

	resp, err := w.manager.Download(ctx, req)
	if err != nil {
		return nil, err
	}

	var parsed *parser.Feed
	if plugin != nil {
		parsed, err = plugin.Parse(ctx, resp)
	}
	if parsed == nil {
		parsed, err = w.manager.Parse(ctx, resp)
	}
	if err != nil {
		return nil, err
	}

	if plugin != nil {
		err = plugin.Postprocess(ctx, parsed)
	} else {
		err = w.manager.Postprocess(ctx, parsed)
	}
	if err != nil {
		return nil, err
	}

	return w.manager.Save(ctx, parsed)
}
