-- name: ListCollections :many
SELECT
  *
FROM
  collections
WHERE
  CASE WHEN sqlc.arg('filter_by_parent_id') THEN
    parent_id = sqlc.narg('parent_id')
  ELSE
    TRUE
  END
ORDER BY
  title ASC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountCollections :one
SELECT
  COUNT(*) AS count
FROM
  collections
WHERE
  CASE WHEN sqlc.arg('filter_by_parent_id') THEN
    parent_id = sqlc.narg('parent_id')
  ELSE
    TRUE
  END;

-- name: InsertCollection :one
INSERT INTO collections(title, parent_id)
  VALUES (sqlc.arg('title'), sqlc.narg('parent_id'))
ON CONFLICT (title)
  DO NOTHING
RETURNING
  *;

-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = sqlc.arg('id');

