package database

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

func Rollback(ctx context.Context, tx pgx.Tx) {
	err := tx.Rollback(ctx)
	if !errors.Is(err, pgx.ErrTxClosed) {
		log.Println(err)
	}
}
