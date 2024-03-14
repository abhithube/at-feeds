package migrations

import (
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

//go:embed postgresql/*
var fs embed.FS

func Migrate(pool *pgxpool.Pool) error {
	sourceInstance, err := iofs.New(fs, "postgresql")
	if err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(pool)
	databaseInstance, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", sourceInstance, "main", databaseInstance)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}

		return err
	}

	return nil
}
