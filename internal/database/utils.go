package database

import (
	"database/sql"
	"errors"
	"log"
)

func Rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if !errors.Is(err, sql.ErrTxDone) {
		log.Println(err)
	}
}
