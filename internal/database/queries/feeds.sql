-- name: ListFeeds :many
SELECT
  *,
(
    SELECT
      count(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = feeds.id
      AND fe.has_read = FALSE) AS unreadCount
FROM
  feeds
ORDER BY
  title ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountFeeds :one
SELECT
  COUNT(*) AS count
FROM
  feeds;

-- name: GetFeed :one
SELECT
  *,
(
    SELECT
      COUNT(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = sqlc.arg('id')) AS entryCount,
(
    SELECT
      count(*)
    FROM
      feed_entries fe
    WHERE
      fe.feed_id = sqlc.arg('id')
      AND fe.has_read = FALSE) AS unreadCount
FROM
  feeds
WHERE
  feeds.id = sqlc.arg('id');

-- name: UpsertFeed :one
INSERT INTO feeds(url, link, title)
  VALUES (sqlc.arg('url'), sqlc.arg('link'), sqlc.arg('title'))
ON CONFLICT (link)
  DO UPDATE SET
    url = excluded.url, title = excluded.title
  RETURNING
    *;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = sqlc.arg('id');

