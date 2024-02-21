package handler

import (
	"database/sql"

	"github.com/abhithube/at-feeds/internal/backup"
	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/task"
)

type Handler struct {
	db            *sql.DB
	queries       *database.Queries
	worker        *task.Worker
	backupManager backup.BackupManager
}

func New(
	db *sql.DB,
	queries *database.Queries,
	worker *task.Worker,
	backupManager backup.BackupManager,
) *Handler {
	return &Handler{
		db:            db,
		queries:       queries,
		worker:        worker,
		backupManager: backupManager,
	}
}
