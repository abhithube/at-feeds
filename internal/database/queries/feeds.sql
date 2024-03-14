-- name: ListFeeds :many
SELECT
  *,
  count(*) OVER () AS total_count,
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

