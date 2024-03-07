// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: collections.sql

package database

import (
	"context"
	"database/sql"
)

const deleteCollection = `-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = ?1
`

func (q *Queries) DeleteCollection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCollection, id)
	return err
}

const insertCollection = `-- name: InsertCollection :one
INSERT INTO collections(title, parent_id)
  VALUES (?1, ?2)
ON CONFLICT (title)
  DO NOTHING
RETURNING
  id, title, parent_id
`

type InsertCollectionParams struct {
	Title    string
	ParentID sql.NullInt64
}

func (q *Queries) InsertCollection(ctx context.Context, arg InsertCollectionParams) (Collection, error) {
	row := q.db.QueryRowContext(ctx, insertCollection, arg.Title, arg.ParentID)
	var i Collection
	err := row.Scan(&i.ID, &i.Title, &i.ParentID)
	return i, err
}