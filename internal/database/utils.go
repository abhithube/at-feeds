package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

func Rollback(ctx context.Context, tx pgx.Tx) {
	err := tx.Rollback(ctx)
	if !errors.Is(err, sql.ErrTxDone) {
		log.Println(err)
	}
}
