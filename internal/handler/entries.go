package handler

import (
	"context"
	"database/sql"
	"time"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListFeedEntries(ctx context.Context, request api.ListFeedEntriesRequestObject) (api.ListFeedEntriesResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	feedID := request.Params.FeedId
	hasRead := request.Params.HasRead
	limit := request.Params.Limit
	page := request.Params.Page

	params := database.ListFeedEntriesParams{Limit: -1}
	if feedID != nil {
		params.FilterByFeedID = true
		params.FeedID = int64(*feedID)
	}
	if hasRead != nil {
		var hasReadInt int64
		if *hasRead {
			hasReadInt = 1
		}
		params.FilterByHasRead = true
		params.HasRead = hasReadInt
	}
	if limit != nil {
		params.Limit = int64(*limit)
		if page != nil {
			params.Offset = (int64(*page) - 1) * params.Limit
		}
	}

	result, err := qtx.ListFeedEntries(ctx, params)
	if err != nil {
		return nil, err
	}

	params2 := database.CountFeedEntriesParams{
		FilterByFeedID:  params.FilterByFeedID,
		FeedID:          params.FeedID,
		FilterByHasRead: params.FilterByHasRead,
		HasRead:         params.HasRead,
	}

	count, err := qtx.CountFeedEntries(ctx, params2)
	if err != nil {
		return nil, err
	}

	data := make([]api.FeedEntry, len(result))
	for i, item := range result {
		publishedAt, err := time.Parse(time.RFC3339, item.PublishedAt)
		if err != nil {
			return nil, err
		}

		entry := api.FeedEntry{
			Id:          int(item.ID),
			Link:        item.Link,
			Title:       item.Title,
			PublishedAt: publishedAt,
			HasRead:     item.HasRead == 1,
			FeedId:      int(item.FeedID),
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

	response := api.ListFeedEntries200JSONResponse{
		Data:    data,
		HasMore: (params.Limit + params.Offset) < count,
	}

	return response, tx.Commit()
}

func (h *Handler) UpdateFeedEntry(ctx context.Context, request api.UpdateFeedEntryRequestObject) (api.UpdateFeedEntryResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	hasRead := request.Body.HasRead

	var hasReadInt int64
	if hasRead != nil && *hasRead {
		hasReadInt = 1
	}
	params2 := database.UpdateFeedEntryParams{
		FeedID:  int64(request.FeedId),
		EntryID: int64(request.EntryId),
		HasRead: sql.NullInt64{Int64: hasReadInt, Valid: hasRead != nil},
	}

	result, err := qtx.UpdateFeedEntry(ctx, params2)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.UpdateFeedEntry404JSONResponse{Message: "Entry not found"}, nil
		}

		return nil, err
	}

	entryResult, err := qtx.GetEntry(ctx, result.EntryID)
	if err != nil {
		return nil, err
	}

	publishedAt, err := time.Parse(time.RFC3339, entryResult.PublishedAt)
	if err != nil {
		return nil, err
	}

	res := api.UpdateFeedEntry200JSONResponse{
		Id:          int(entryResult.ID),
		Link:        entryResult.Link,
		Title:       entryResult.Title,
		PublishedAt: publishedAt,
		HasRead:     result.HasRead == 1,
		FeedId:      request.FeedId,
	}
	if entryResult.Author.Valid {
		res.Author = &entryResult.Author.String
	}
	if entryResult.Content.Valid {
		res.Content = &entryResult.Content.String
	}
	if entryResult.ThumbnailUrl.Valid {
		res.ThumbnailUrl = &entryResult.ThumbnailUrl.String
	}

	return res, tx.Commit()
}
