// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: entries.sql

package database

import (
	"context"
	"database/sql"
)

const countEntries = `-- name: CountEntries :one
SELECT
  COUNT(*) AS count
FROM
  entries
WHERE
  CASE WHEN ?1 THEN
    feed_id = ?2
  ELSE
    TRUE
  END
  AND CASE WHEN ?3 THEN
    has_read = ?4
  ELSE
    TRUE
  END
`

type CountEntriesParams struct {
	FilterByFeedID  interface{}
	FeedID          int64
	FilterByHasRead interface{}
	HasRead         int64
}

func (q *Queries) CountEntries(ctx context.Context, arg CountEntriesParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countEntries,
		arg.FilterByFeedID,
		arg.FeedID,
		arg.FilterByHasRead,
		arg.HasRead,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getEntry = `-- name: GetEntry :one
SELECT
  id, link, title, published_at, author, content, thumbnail_url, has_read, feed_id
FROM
  entries
WHERE
  id = ?1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.Title,
		&i.PublishedAt,
		&i.Author,
		&i.Content,
		&i.ThumbnailUrl,
		&i.HasRead,
		&i.FeedID,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT
  id, link, title, published_at, author, content, thumbnail_url, has_read, feed_id
FROM
  entries
WHERE
  CASE WHEN ?1 THEN
    feed_id = ?2
  ELSE
    TRUE
  END
  AND CASE WHEN ?3 THEN
    has_read = ?4
  ELSE
    TRUE
  END
ORDER BY
  published_at DESC
LIMIT ?6 OFFSET ?5
`

type ListEntriesParams struct {
	FilterByFeedID  interface{}
	FeedID          int64
	FilterByHasRead interface{}
	HasRead         int64
	Offset          int64
	Limit           int64
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries,
		arg.FilterByFeedID,
		arg.FeedID,
		arg.FilterByHasRead,
		arg.HasRead,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Link,
			&i.Title,
			&i.PublishedAt,
			&i.Author,
			&i.Content,
			&i.ThumbnailUrl,
			&i.HasRead,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntry = `-- name: UpdateEntry :one
UPDATE
  entries
SET
  has_read = ?1
WHERE
  id = ?2
RETURNING
  id, link, title, published_at, author, content, thumbnail_url, has_read, feed_id
`

type UpdateEntryParams struct {
	HasRead int64
	ID      int64
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntry, arg.HasRead, arg.ID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Link,
		&i.Title,
		&i.PublishedAt,
		&i.Author,
		&i.Content,
		&i.ThumbnailUrl,
		&i.HasRead,
		&i.FeedID,
	)
	return i, err
}

const upsertEntry = `-- name: UpsertEntry :exec
INSERT INTO entries(link, title, published_at, author, content, thumbnail_url, feed_id)
  VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7)
ON CONFLICT (feed_id, link)
  DO UPDATE SET
    title = excluded.title, published_at = excluded.published_at, author = excluded.author, content = excluded.content, thumbnail_url = excluded.thumbnail_url
`

type UpsertEntryParams struct {
	Link         string
	Title        string
	PublishedAt  string
	Author       sql.NullString
	Content      sql.NullString
	ThumbnailUrl sql.NullString
	FeedID       int64
}

func (q *Queries) UpsertEntry(ctx context.Context, arg UpsertEntryParams) error {
	_, err := q.db.ExecContext(ctx, upsertEntry,
		arg.Link,
		arg.Title,
		arg.PublishedAt,
		arg.Author,
		arg.Content,
		arg.ThumbnailUrl,
		arg.FeedID,
	)
	return err
}
