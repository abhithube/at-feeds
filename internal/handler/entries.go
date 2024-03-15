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
		data[i] = mapToAPIFeedEntry(item.Entry, item.FeedEntry)
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

	result1, err := h.queries.UpdateFeedEntry(ctx, params2)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.UpdateFeedEntry404JSONResponse{Message: "Entry not found"}, nil
		}

		return nil, err
	}

	result2, err := h.queries.GetEntry(ctx, result1.EntryID)
	if err != nil {
		return nil, err
	}

	res := api.UpdateFeedEntry200JSONResponse{
		Data: mapToAPIFeedEntry(result2, result1),
	}

	return res, nil
}
