package task

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/parser"
	"golang.org/x/net/html"
)

type Manager struct {
	db         *sql.DB
	queries    *database.Queries
	httpClient *http.Client
}

func NewManager(db *sql.DB, queries *database.Queries, httpClient *http.Client) *Manager {
	return &Manager{
		db:         db,
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
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := m.queries.WithTx(tx)

	var url sql.NullString
	if feed.URL != "" {
		url = sql.NullString{String: feed.URL, Valid: true}
	}

	params := database.UpsertFeedParams{
		Url:   url,
		Link:  feed.Link,
		Title: feed.Title,
	}
	inserted, err := qtx.UpsertFeed(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, entry := range feed.Entries {
		var author, content, thumbnailURL sql.NullString
		if entry.Author != "" {
			author = sql.NullString{String: entry.Author, Valid: true}
		}
		if entry.Content != "" {
			content = sql.NullString{String: entry.Content, Valid: true}
		}
		if entry.ThumbnailURL != "" {
			thumbnailURL = sql.NullString{String: entry.ThumbnailURL, Valid: true}
		}
		publishedAt, err := entry.PublishedAt.MarshalText()
		if err != nil {
			return nil, err
		}
		params := database.UpsertEntryParams{
			Link:         entry.Link,
			Title:        entry.Title,
			PublishedAt:  string(publishedAt),
			Author:       author,
			Content:      content,
			ThumbnailUrl: thumbnailURL,
		}

		insertedEntry, err := qtx.UpsertEntry(ctx, params)
		if err != nil {
			return nil, err
		}

		params2 := database.UpsertFeedEntryParams{
			EntryID: insertedEntry.ID,
			FeedID:  inserted.ID,
		}

		if err = qtx.UpsertFeedEntry(ctx, params2); err != nil {
			return nil, err
		}
	}

	return &inserted, tx.Commit()
}
