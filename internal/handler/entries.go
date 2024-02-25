package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListEntries(ctx context.Context, request api.ListEntriesRequestObject) (api.ListEntriesResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	params := database.ListEntriesParams{Limit: -1}
	if request.Params.FeedId != nil {
		params.FilterByFeedID = true
		params.FeedID = int64(*request.Params.FeedId)
	}
	if request.Params.HasRead != nil {
		var hasRead int64
		if *request.Params.HasRead {
			hasRead = 1
		}
		params.FilterByHasRead = true
		params.HasRead = hasRead
	}
	if request.Params.Limit != nil {
		params.Limit = int64(*request.Params.Limit)
		if request.Params.Page != nil {
			params.Offset = (int64(*request.Params.Page) - 1) * params.Limit
		}
	}

	result, err := qtx.ListEntries(ctx, params)
	if err != nil {
		return nil, err
	}

	params2 := database.CountEntriesParams{
		FilterByFeedID:  params.FilterByFeedID,
		FeedID:          params.FeedID,
		FilterByHasRead: params.FilterByHasRead,
		HasRead:         params.HasRead,
	}

	count, err := qtx.CountEntries(ctx, params2)
	if err != nil {
		return nil, err
	}

	data := make([]api.Entry, len(result))
	for i, item := range result {
		publishedAt, err := time.Parse(time.RFC3339, item.PublishedAt)
		if err != nil {
			return nil, err
		}

		entry := api.Entry{
			Id:          int(item.ID),
			Link:        item.Link,
			Title:       item.Title,
			PublishedAt: publishedAt,
			HasRead:     item.HasRead == 1,
		}
		if item.Author.Valid {
			entry.Author = &item.Author.String
		}
		if item.Content.Valid {
			entry.Content = &item.Content.String
		}
		if item.ThumbnailUrl.Valid {
			entry.ThumbnailUrl = &item.ThumbnailUrl.String
		}

		data[i] = entry
	}

	response := api.ListEntries200JSONResponse{
		Data:    data,
		HasMore: (params.Limit + params.Offset) < count,
	}

	return response, tx.Commit()
}

func (h *Handler) UpdateEntry(ctx context.Context, request api.UpdateEntryRequestObject) (api.UpdateEntryResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	if _, err := qtx.GetEntry(ctx, int64(request.Id)); err != nil {
		if err == sql.ErrNoRows {
			return api.UpdateEntry404JSONResponse{Message: "Entry not found"}, nil
		}

		return nil, err
	}

	var hasRead int64
	if request.Body.HasRead {
		hasRead = 1
	}
	params := database.UpdateEntryParams{
		ID:      int64(request.Id),
		HasRead: hasRead,
	}

	result, err := qtx.UpdateEntry(ctx, params)
	if err != nil {
		return nil, err
	}

	publishedAt, err := time.Parse(time.RFC3339, result.PublishedAt)
	if err != nil {
		return nil, err
	}

	res := api.UpdateEntry200JSONResponse{
		Id:          int(result.ID),
		Link:        result.Link,
		Title:       result.Title,
		PublishedAt: publishedAt,
		HasRead:     result.HasRead == 1,
	}
	if result.Author.Valid {
		res.Author = &result.Author.String
	}
	if result.Content.Valid {
		res.Content = &result.Content.String
	}
	if result.ThumbnailUrl.Valid {
		res.ThumbnailUrl = &result.ThumbnailUrl.String
	}

	return res, tx.Commit()
}
