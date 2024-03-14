package task

import (
	"context"
	"log"

	"github.com/abhithube/at-feeds/internal/database"
)

type RefreshJob struct {
	queries *database.Queries
	worker  *Worker
}

func NewRefreshJob(queries *database.Queries, worker *Worker) *RefreshJob {
	return &RefreshJob{queries: queries, worker: worker}
}

func (j *RefreshJob) Run(ctx context.Context) {
	feeds, err := j.queries.ListFeeds(ctx, database.ListFeedsParams{})
	if err != nil {
		log.Fatal(err)
	}

	urls := make([]string, len(feeds))
	for i, feed := range feeds {
		urls[i] = feed.Link
		if feed.Url.Valid {
			urls[i] = feed.Url.String
		}
	}

	if err := j.worker.RunAll(ctx, urls); err != nil {
		log.Fatal(err)
	}
}
