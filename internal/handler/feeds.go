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
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()
	qtx := h.queries.WithTx(tx)

	params := database.ListFeedsParams{Limit: -1}
	if request.Params.Limit != nil {
		params.Limit = int64(*request.Params.Limit)
		params.Offset = int64(*request.Params.Page-1) * params.Limit
	}

	result, err := qtx.ListFeeds(ctx, params)
	if err != nil {
		return nil, err
	}

	count, err := qtx.CountFeeds(ctx)
	if err != nil {
		return nil, err
	}

	arr := make([]api.Feed, len(result))
	for i, item := range result {
		feed := api.Feed{
			Id:          int(item.ID),
			Link:        item.Link,
			Title:       item.Title,
			UnreadCount: int(item.Unreadcount),
		}
		if item.Url.Valid {
			feed.Url = &item.Url.String
		}

		arr[i] = feed
	}

	response := api.ListFeeds200JSONResponse{
		Data:    arr,
		HasMore: (params.Limit + params.Offset) < count,
	}

	return response, tx.Commit()
}

func (h *Handler) GetFeed(ctx context.Context, request api.GetFeedRequestObject) (api.GetFeedResponseObject, error) {
	result, err := h.queries.GetFeed(ctx, int64(request.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return api.GetFeed404JSONResponse{Message: "Feed not found"}, nil
		} else {
			return nil, err
		}
	}

	entryCount := int(result.Entrycount)
	response := api.GetFeed200JSONResponse{
		Id:          int(result.ID),
		Link:        result.Link,
		Title:       result.Title,
		EntryCount:  &entryCount,
		UnreadCount: int(result.Unreadcount),
	}
	if result.Url.Valid {
		response.Url = &result.Url.String
	}

	return response, nil
}

func (h *Handler) CreateFeed(ctx context.Context, request api.CreateFeedRequestObject) (api.CreateFeedResponseObject, error) {
	handleError := func(err error) api.CreateFeed400JSONResponse {
		return api.CreateFeed400JSONResponse{Message: err.Error()}
	}

	result, err := h.worker.Run(ctx, request.Body.Url)
	if err != nil {
		return handleError(err), nil
	}

	response := api.CreateFeed201JSONResponse{
		Id:    int(result.ID),
		Link:  result.Link,
		Title: result.Title,
	}
	if result.Url.Valid {
		response.Url = &result.Url.String
	}

	return response, nil
}

func (h *Handler) DeleteFeed(ctx context.Context, request api.DeleteFeedRequestObject) (api.DeleteFeedResponseObject, error) {
	err := h.queries.DeleteFeed(ctx, int64(request.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return api.DeleteFeed404JSONResponse{Message: "Feed not found"}, nil
		} else {
			return nil, err
		}
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

	result, err := h.queries.ListFeeds(ctx, database.ListFeedsParams{Limit: -1})
	if err != nil {
		return handleError(err), nil
	}

	feedMap := make(map[string]struct{})
	for _, feed := range result {
		if feed.Url.Valid {
			feedMap[feed.Url.String] = struct{}{}
		}
	}

	feedURLs := make([]string, 0)

	for _, item := range items {
		if _, exists := feedMap[item.URL]; !exists {
			feedURLs = append(feedURLs, item.URL)
		}
	}

	h.worker.RunAll(ctx, feedURLs)

	return api.ImportFeeds200Response{}, nil
}

func (h *Handler) ExportFeeds(ctx context.Context, request api.ExportFeedsRequestObject) (api.ExportFeedsResponseObject, error) {
	result, err := h.queries.ListFeeds(ctx, database.ListFeedsParams{Limit: -1})
	if err != nil {
		return api.ExportFeeds500JSONResponse{Message: err.Error()}, nil
	}

	items := make([]backup.BackupItem, len(result))
	for i, feed := range result {
		if !feed.Url.Valid {
			continue
		}

		item := backup.BackupItem{
			URL:   feed.Url.String,
			Link:  feed.Link,
			Title: feed.Title,
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
