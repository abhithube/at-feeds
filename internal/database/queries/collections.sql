-- name: ListCollections :many
SELECT
  *,
  count(*) OVER () AS total_count
FROM
  collections
WHERE
  CASE WHEN sqlc.arg('filter_by_parent_id') THEN
    CASE WHEN sqlc.arg('parent_id') < 0 THEN
      parent_id IS NULL
    ELSE
      parent_id = sqlc.arg('parent_id')
    END
  ELSE
    TRUE
  END
ORDER BY
  title ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: GetCollection :one
SELECT
  *
FROM
  collections
WHERE
  id = sqlc.arg('id');

-- name: InsertCollection :one
INSERT INTO collections(title, parent_id)
  VALUES (sqlc.arg('title'), sqlc.narg('parent_id'))
ON CONFLICT (title, parent_id)
  DO NOTHING
RETURNING
  *;

-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = sqlc.arg('id');

