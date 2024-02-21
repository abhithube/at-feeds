-- name: ListEntries :many
SELECT
  *
FROM
  entries
WHERE
  CASE WHEN sqlc.arg('filter_by_feed_id') THEN
    feed_id = sqlc.arg('feed_id')
  ELSE
    TRUE
  END
  AND CASE WHEN sqlc.arg('filter_by_has_read') THEN
    has_read = sqlc.arg('has_read')
  ELSE
    TRUE
  END
ORDER BY
  published_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountEntries :one
SELECT
  COUNT(*) AS count
FROM
  entries
WHERE
  CASE WHEN sqlc.arg('filter_by_feed_id') THEN
    feed_id = sqlc.arg('feed_id')
  ELSE
    TRUE
  END
  AND CASE WHEN sqlc.arg('filter_by_has_read') THEN
    has_read = sqlc.arg('has_read')
  ELSE
    TRUE
  END;

-- name: GetEntry :one
SELECT
  *
FROM
  entries
WHERE
  id = sqlc.arg('id');

-- name: UpsertEntry :exec
INSERT INTO entries(link, title, published_at, author, content, thumbnail_url, feed_id)
  VALUES (sqlc.arg('link'), sqlc.arg('title'), sqlc.arg('published_at'), sqlc.arg('author'), sqlc.arg('content'), sqlc.arg('thumbnail_url'), sqlc.arg('feed_id'))
ON CONFLICT (feed_id, link)
  DO UPDATE SET
    title = excluded.title, published_at = excluded.published_at, author = excluded.author, content = excluded.content, thumbnail_url = excluded.thumbnail_url;

-- name: UpdateEntry :one
UPDATE
  entries
SET
  has_read = sqlc.arg('has_read')
WHERE
  id = sqlc.arg('id')
RETURNING
  *;

