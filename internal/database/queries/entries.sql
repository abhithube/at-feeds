-- name: GetEntry :one
SELECT
  *
FROM
  entries
WHERE
  id = sqlc.arg('id');

-- name: CreateEntry :one
INSERT INTO entries(link, title, published_at, author, content, thumbnail_url)
  VALUES (sqlc.arg('link'), sqlc.arg('title'), sqlc.arg('published_at'), sqlc.arg('author'), sqlc.arg('content'), sqlc.arg('thumbnail_url'))
ON CONFLICT (link)
  DO UPDATE SET
    title = excluded.title, published_at = excluded.published_at, author = excluded.author, content = excluded.content, thumbnail_url = excluded.thumbnail_url
  RETURNING
    *;

