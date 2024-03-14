package task

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/parser"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/html"
)

type Manager struct {
	pool       *pgxpool.Pool
	queries    *database.Queries
	httpClient *http.Client
}

func NewManager(pool *pgxpool.Pool, queries *database.Queries, httpClient *http.Client) *Manager {
	return &Manager{
		pool:       pool,
		httpClient: httpClient,
		queries:    queries,
	}
}

func (m *Manager) Preprocess(ctx context.Context, feedURL *url.URL) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, http.MethodGet, feedURL.String(), nil)
}

func (m *Manager) Download(_ context.Context, req *http.Request) (*http.Response, error) {
	return m.httpClient.Do(req)
}

func (m *Manager) Parse(_ context.Context, resp *http.Response) (*parser.Feed, error) {
	defer resp.Body.Close()

	parsedURL := resp.Request.URL
	feedURL := parsedURL.String()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if parser.IsHTMLDocument(resp.Header, data) {
		doc, err := html.Parse(bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}

		n := parser.Query(doc, "link[type=\"application/rss+xml\"]")
		if n == nil {
			return nil, errors.New("URL not supported")
		}

		feedURL = parser.Attr(n, "href")
		resp, err := m.httpClient.Get(feedURL)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	}

	if parser.IsAtomFeed(resp.Header, data) {
		feed, err := parser.NewAtomParser().Parse(data)
		if err != nil {
			return nil, err
		}

		feed.URL = feedURL
		return feed, nil
	}

	if parser.IsRSSFeed(resp.Header, data) {
		feed, err := parser.NewRSSParser().Parse(data)
		if err != nil {
			return nil, err
		}

		feed.URL = feedURL
		return feed, nil
	}

	msg := fmt.Sprintf("no parser found for %s", feedURL)
	return nil, errors.New(msg)
}

func (m *Manager) Postprocess(_ context.Context, feed *parser.Feed) error {
	if len(feed.Entries) > 0 && feed.Entries[0].ThumbnailURL == "" {
		return parser.NewThumbnailDescriptionProcessor().Postprocess(feed)
	}

	return nil
}

func (m *Manager) Save(ctx context.Context, feed *parser.Feed) (*database.Feed, error) {
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	defer database.Rollback(ctx, tx)

	qtx := m.queries.WithTx(tx)

	params := database.UpsertFeedParams{
		Link:  feed.Link,
		Title: feed.Title,
	}
	params.Url.String = feed.URL
	params.Url.Valid = feed.URL != ""

	inserted, err := qtx.UpsertFeed(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, entry := range feed.Entries {
		params := database.UpsertEntryParams{
			Link:  entry.Link,
			Title: entry.Title,
		}
		params.ThumbnailUrl.String = entry.ThumbnailURL
		params.ThumbnailUrl.Valid = entry.ThumbnailURL != ""
		params.PublishedAt.Time = entry.PublishedAt
		params.PublishedAt.Valid = true
		params.Content.String = entry.Content
		params.Content.Valid = entry.Content != ""
		params.Author.String = entry.Author
		params.Author.Valid = entry.Author != ""

		insertedEntry, err := qtx.UpsertEntry(ctx, params)
		if err != nil {
			return nil, err
		}

		params2 := database.UpsertFeedEntryParams{
			EntryID: int32(insertedEntry.ID),
			FeedID:  int32(inserted.ID),
		}

		if err = qtx.UpsertFeedEntry(ctx, params2); err != nil {
			return nil, err
		}
	}

	return &inserted, tx.Commit(ctx)
}
