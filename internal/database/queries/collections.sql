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

