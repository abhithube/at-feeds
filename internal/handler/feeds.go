package handler

import (
	"bytes"
	"context"
	"database/sql"
	"io"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/backup"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListFeeds(ctx context.Context, request api.ListFeedsRequestObject) (api.ListFeedsResponseObject, error) {
	collectionID := request.Params.CollectionId
	limit := request.Params.Limit
	page := request.Params.Page

	params := database.ListFeedsParams{}
	if limit != nil && *limit >= 0 {
		params.Limit.Int32 = int32(*limit)
		params.Limit.Valid = true
		if page != nil {
			params.Offset = (int32(*page) - 1) * params.Limit.Int32
		}
	}
	if collectionID != nil {
		params.FilterByCollectionID = true
		params.CollectionID = collectionID
	}

	result, err := h.queries.ListFeeds(ctx, params)
	if err != nil {
		return nil, err
	}

	arr := make([]api.Feed, len(result))
	for i, item := range result {
		arr[i] = mapToAPIFeed(item.Feed, &item.TotalEntryCount, &item.UnreadEntryCount)
	}

	var hasMore bool
	if len(result) > 0 {
		hasMore = int64(params.Offset+params.Limit.Int32) < result[0].TotalCount
	}
	response := api.ListFeeds200JSONResponse{
		Data:    arr,
		HasMore: hasMore,
	}

	return response, nil
}

func (h *Handler) GetFeed(ctx context.Context, request api.GetFeedRequestObject) (api.GetFeedResponseObject, error) {
	result, err := h.queries.GetFeed(ctx, int32(request.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return api.GetFeed404JSONResponse{Message: "Feed not found"}, nil
		}

		return nil, err
	}

	response := api.GetFeed200JSONResponse{
		Data: mapToAPIFeed(result.Feed, &result.TotalEntryCount, &result.UnreadEntryCount),
	}

	return response, nil
}

func (h *Handler) CreateFeed(ctx context.Context, request api.CreateFeedRequestObject) (api.CreateFeedResponseObject, error) {
	result, err := h.worker.Run(ctx, request.Body.Url)
	if err != nil {
		return api.CreateFeed400JSONResponse{Message: err.Error()}, nil
	}

	response := api.CreateFeed201JSONResponse{
		Data: mapToAPIFeed(*result, nil, nil),
	}

	return response, nil
}

func (h *Handler) UpdateFeed(ctx context.Context, request api.UpdateFeedRequestObject) (api.UpdateFeedResponseObject, error) {
	collectionID := request.Body.CollectionId

	params := database.UpdateFeedParams{
		ID: int32(request.Id),
	}
	if collectionID != nil {
		params.CollectionID.Int32 = int32(*collectionID)
		params.CollectionID.Valid = true
	}

	result, err := h.queries.UpdateFeed(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return api.UpdateFeed404JSONResponse{Message: "Feed not found"}, nil
		}

		return nil, err
	}

	response := api.UpdateFeed200JSONResponse{
		Data: mapToAPIFeed(result, nil, nil),
	}

	return response, nil
}

func (h *Handler) DeleteFeed(ctx context.Context, request api.DeleteFeedRequestObject) (api.DeleteFeedResponseObject, error) {
	err := h.queries.DeleteFeed(ctx, int32(request.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return api.DeleteFeed404JSONResponse{Message: "Feed not found"}, nil
		}

		return nil, err
	}

	return api.DeleteFeed204Response{}, nil
}

func (h *Handler) ImportFeeds(ctx context.Context, request api.ImportFeedsRequestObject) (api.ImportFeedsResponseObject, error) {
	handleError := func(err error) api.ImportFeedsResponseObject {
		return api.ImportFeeds500JSONResponse{Message: err.Error()}
	}

	data, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	items, err := h.backupManager.Import(data)
	if err != nil {
		return handleError(err), nil
	}

	result, err := h.queries.ListFeeds(ctx, database.ListFeedsParams{})
	if err != nil {
		return handleError(err), nil
	}

	feedMap := make(map[string]struct{})
	for _, item := range result {
		if item.Feed.Url.Valid {
			feedMap[item.Feed.Url.String] = struct{}{}
		}
	}

	feedURLs := make([]string, 0)

	for _, item := range items {
		if _, exists := feedMap[item.URL]; !exists {
			feedURLs = append(feedURLs, item.URL)
		}
	}

	if err = h.worker.RunAll(ctx, feedURLs); err != nil {
		return handleError(err), nil
	}

	return api.ImportFeeds200Response{}, nil
}

func (h *Handler) ExportFeeds(ctx context.Context, _ api.ExportFeedsRequestObject) (api.ExportFeedsResponseObject, error) {
	result, err := h.queries.ListFeeds(ctx, database.ListFeedsParams{})
	if err != nil {
		return api.ExportFeeds500JSONResponse{Message: err.Error()}, nil
	}

	items := make([]backup.Item, len(result))
	for i, item := range result {
		if !item.Feed.Url.Valid {
			continue
		}

		item := backup.Item{
			URL:   item.Feed.Url.String,
			Link:  item.Feed.Link,
			Title: item.Feed.Title,
		}
		items[i] = item
	}

	data, err := h.backupManager.Export(items)
	if err != nil {
		return api.ExportFeeds500JSONResponse{Message: err.Error()}, nil
	}

	body := bytes.NewReader(data)
	response := api.ExportFeeds200ApplicationoctetStreamResponse{
		Body:          body,
		ContentLength: body.Size(),
	}

	return response, nil
}
