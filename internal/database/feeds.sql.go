// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: feeds.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds(url, link, title)
  VALUES ($1, $2, $3)
ON CONFLICT (link)
  DO UPDATE SET
    url = excluded.url, title = excluded.title
  RETURNING
    id, url, link, title, collection_id
`

type CreateFeedParams struct {
	Url   pgtype.Text
	Link  string
	Title string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, createFeed, arg.Url, arg.Link, arg.Title)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Link,
		&i.Title,
		&i.CollectionID,
	)
	return i, err
}

const deleteFeed = `-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = $1
`

func (q *Queries) DeleteFeed(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteFeed, id)
	return err
}

const getFeed = `-- name: GetFeed :one
SELECT
  f.id, f.url, f.link, f.title, f.collection_id,
(
    SELECT
      COUNT(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = $1) AS total_entry_count,
(
    SELECT
      count(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = $1
      AND fe.has_read = FALSE) AS unread_entry_count
FROM
  feeds f
WHERE
  f.id = $1
`

type GetFeedRow struct {
	Feed             Feed
	TotalEntryCount  int64
	UnreadEntryCount int64
}

func (q *Queries) GetFeed(ctx context.Context, id int32) (GetFeedRow, error) {
	row := q.db.QueryRow(ctx, getFeed, id)
	var i GetFeedRow
	err := row.Scan(
		&i.Feed.ID,
		&i.Feed.Url,
		&i.Feed.Link,
		&i.Feed.Title,
		&i.Feed.CollectionID,
		&i.TotalEntryCount,
		&i.UnreadEntryCount,
	)
	return i, err
}

const listFeeds = `-- name: ListFeeds :many
SELECT
  f.id, f.url, f.link, f.title, f.collection_id,
  count(*) OVER () AS total_count,
(
    SELECT
      COUNT(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = $1) AS total_entry_count,
(
    SELECT
      count(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = f.id
      AND fe.has_read = FALSE) AS unread_entry_count
FROM
  feeds f
WHERE
  CASE WHEN $2 THEN
    CASE WHEN $3 < 0 THEN
      collection_id IS NULL
    ELSE
      collection_id = $3
    END
  ELSE
    TRUE
  END
ORDER BY
  title ASC
LIMIT $5 OFFSET $4
`

type ListFeedsParams struct {
	ID                   int32
	FilterByCollectionID interface{}
	CollectionID         interface{}
	Offset               int32
	Limit                pgtype.Int4
}

type ListFeedsRow struct {
	Feed             Feed
	TotalCount       int64
	TotalEntryCount  int64
	UnreadEntryCount int64
}

func (q *Queries) ListFeeds(ctx context.Context, arg ListFeedsParams) ([]ListFeedsRow, error) {
	rows, err := q.db.Query(ctx, listFeeds,
		arg.ID,
		arg.FilterByCollectionID,
		arg.CollectionID,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFeedsRow
	for rows.Next() {
		var i ListFeedsRow
		if err := rows.Scan(
			&i.Feed.ID,
			&i.Feed.Url,
			&i.Feed.Link,
			&i.Feed.Title,
			&i.Feed.CollectionID,
			&i.TotalCount,
			&i.TotalEntryCount,
			&i.UnreadEntryCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFeed = `-- name: UpdateFeed :one
UPDATE
  feeds
SET
  collection_id = coalesce($1, collection_id)
WHERE
  id = $2
RETURNING
  id, url, link, title, collection_id
`

type UpdateFeedParams struct {
	CollectionID pgtype.Int4
	ID           int32
}

func (q *Queries) UpdateFeed(ctx context.Context, arg UpdateFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, updateFeed, arg.CollectionID, arg.ID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Link,
		&i.Title,
		&i.CollectionID,
	)
	return i, err
}
