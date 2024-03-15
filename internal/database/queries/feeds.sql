-- name: ListFeeds :many
SELECT
  sqlc.embed(f),
  count(*) OVER () AS total_count,
(
    SELECT
      COUNT(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = sqlc.arg('id')) AS total_entry_count,
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
  CASE WHEN sqlc.arg('filter_by_collection_id') THEN
    CASE WHEN sqlc.arg('collection_id') < 0 THEN
      collection_id IS NULL
    ELSE
      collection_id = sqlc.arg('collection_id')
    END
  ELSE
    TRUE
  END
ORDER BY
  title ASC
LIMIT sqlc.narg('limit') OFFSET sqlc.arg('offset');

-- name: GetFeed :one
SELECT
  sqlc.embed(f),
(
    SELECT
      COUNT(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = sqlc.arg('id')) AS total_entry_count,
(
    SELECT
      count(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = sqlc.arg('id')
      AND fe.has_read = FALSE) AS unread_entry_count
FROM
  feeds f
WHERE
  f.id = sqlc.arg('id');

-- name: CreateFeed :one
INSERT INTO feeds(url, link, title)
  VALUES (sqlc.arg('url'), sqlc.arg('link'), sqlc.arg('title'))
ON CONFLICT (link)
  DO UPDATE SET
    url = excluded.url, title = excluded.title
  RETURNING
    *;

-- name: UpdateFeed :one
UPDATE
  feeds
SET
  collection_id = coalesce(sqlc.narg('collection_id'), collection_id)
WHERE
  id = sqlc.arg('id')
RETURNING
  *;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = sqlc.arg('id');

