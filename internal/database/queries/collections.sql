-- name: ListCollections :many
SELECT
  *,
  count(*) OVER () AS total_count
FROM
  collections
ORDER BY
  title ASC
LIMIT sqlc.narg('limit') OFFSET sqlc.arg('offset');

-- name: GetCollection :one
SELECT
  *
FROM
  collections
WHERE
  id = sqlc.arg('id');

-- name: InsertCollection :one
INSERT INTO collections(title)
  VALUES (sqlc.arg('title'))
RETURNING
  *;

-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = sqlc.arg('id');

