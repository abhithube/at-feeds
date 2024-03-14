package handler

import (
	"github.com/abhithube/at-feeds/internal/backup"
	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/task"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	pool          *pgxpool.Pool
	queries       *database.Queries
	worker        *task.Worker
	backupManager backup.Manager
}

func New(
	pool *pgxpool.Pool,
	queries *database.Queries,
	worker *task.Worker,
	backupManager backup.Manager,
) *Handler {
	return &Handler{
		pool:          pool,
		queries:       queries,
		worker:        worker,
		backupManager: backupManager,
	}
}
