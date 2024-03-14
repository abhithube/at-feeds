package handler

import (
	"context"
	"database/sql"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListFeedEntries(ctx context.Context, request api.ListFeedEntriesRequestObject) (api.ListFeedEntriesResponseObject, error) {
	feedID := request.Params.FeedId
	hasRead := request.Params.HasRead
	limit := request.Params.Limit
	page := request.Params.Page

	params := database.ListFeedEntriesParams{}
	if feedID != nil {
		params.FilterByFeedID = true
		params.FeedID = int32(*feedID)
	}
	if hasRead != nil {
		params.FilterByHasRead = true
		params.HasRead = *hasRead
	}
	if limit != nil && *limit >= 0 {
		params.Limit.Int32 = int32(*limit)
		params.Limit.Valid = true
		if page != nil {
			params.Offset = (int32(*page) - 1) * params.Limit.Int32
		}
	}

	result, err := h.queries.ListFeedEntries(ctx, params)
	if err != nil {
		return nil, err
	}

	data := make([]api.FeedEntry, len(result))
	for i, item := range result {
		entry := api.FeedEntry{
			Id:          int(item.ID),
			Link:        item.Link,
			Title:       item.Title,
			PublishedAt: item.PublishedAt.Time,
			HasRead:     item.HasRead,
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

	var hasMore bool
	if len(result) > 0 {
		hasMore = int64(params.Offset+params.Limit.Int32) < result[0].TotalCount
	}
	response := api.ListFeedEntries200JSONResponse{
		Data:    data,
		HasMore: hasMore,
	}

	return response, nil
}

func (h *Handler) UpdateFeedEntry(ctx context.Context, request api.UpdateFeedEntryRequestObject) (api.UpdateFeedEntryResponseObject, error) {
	hasRead := request.Body.HasRead

	params2 := database.UpdateFeedEntryParams{
		FeedID:  int32(request.FeedId),
		EntryID: int32(request.EntryId),
	}
	if hasRead != nil {
		params2.HasRead.Bool = *hasRead
		params2.HasRead.Valid = true
	}

	result, err := h.queries.UpdateFeedEntry(ctx, params2)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.UpdateFeedEntry404JSONResponse{Message: "Entry not found"}, nil
		}

		return nil, err
	}

	entryResult, err := h.queries.GetEntry(ctx, result.EntryID)
	if err != nil {
		return nil, err
	}

	res := api.UpdateFeedEntry200JSONResponse{
		Id:          int(entryResult.ID),
		Link:        entryResult.Link,
		Title:       entryResult.Title,
		PublishedAt: entryResult.PublishedAt.Time,
		HasRead:     result.HasRead,
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

	return res, nil
}
