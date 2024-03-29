-- name: ListFeedEntries :many
SELECT
  sqlc.embed(e),
  sqlc.embed(fe),
  count(*) OVER () AS total_count
FROM
  feed_entries fe
  JOIN entries e ON e.id = fe.entry_id
WHERE
  CASE WHEN sqlc.arg('filter_by_feed_id') THEN
    fe.feed_id = sqlc.arg('feed_id')
  ELSE
    TRUE
  END
  AND CASE WHEN sqlc.arg('filter_by_has_read') THEN
    fe.has_read = sqlc.arg('has_read')
  ELSE
    TRUE
  END
ORDER BY
  e.published_at DESC
LIMIT sqlc.narg('limit') OFFSET sqlc.arg('offset');

-- name: GetFeedEntry :one
SELECT
  *
FROM
  feed_entries
WHERE
  feed_id = sqlc.arg('feed_id')
  AND entry_id = sqlc.arg('entry_id');

-- name: CreateFeedEntry :exec
INSERT INTO feed_entries(entry_id, feed_id)
  VALUES (sqlc.arg('entry_id'), sqlc.arg('feed_id'))
ON CONFLICT (feed_id, entry_id)
  DO NOTHING
RETURNING
  *;

-- name: UpdateFeedEntry :one
UPDATE
  feed_entries
SET
  has_read = coalesce(sqlc.narg('has_read'), has_read)
WHERE
  feed_id = sqlc.arg('feed_id')
  AND entry_id = sqlc.arg('entry_id')
RETURNING
  *;

