package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListCollections(ctx context.Context, request api.ListCollectionsRequestObject) (api.ListCollectionsResponseObject, error) {
	limit := request.Params.Limit
	page := request.Params.Page

	params := database.ListCollectionsParams{Limit: -1}
	if limit != nil {
		params.Limit = int64(*limit)
		if page != nil {
			params.Offset = (int64(*page) - 1) * params.Limit
		}
	}

	result, err := h.queries.ListCollections(ctx, params)
	if err != nil {
		return nil, err
	}

	arr := make([]api.Collection, len(result))
	for i, item := range result {
		collection := api.Collection{
			Id:    int(item.ID),
			Title: item.Title,
		}

		arr[i] = collection
	}

	var hasMore bool
	if len(result) > 0 {
		hasMore = (params.Offset + params.Limit) < result[0].TotalCount
	}
	response := api.ListCollections200JSONResponse{
		Data:    arr,
		HasMore: hasMore,
	}

	return response, nil
}

func (h *Handler) CreateCollection(ctx context.Context, request api.CreateCollectionRequestObject) (api.CreateCollectionResponseObject, error) {
	title := request.Body.Title
	if len(title) == 0 {
		return api.CreateCollection400JSONResponse{Message: "'title' cannot be empty"}, nil
	}

	result, err := h.queries.InsertCollection(ctx, title)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, sql.ErrNoRows) {
			msg = fmt.Sprintf("Collection already exists with title '%s'", title)
		}

		return api.CreateCollection400JSONResponse{Message: msg}, nil
	}

	response := api.CreateCollection201JSONResponse{
		Id:    int(result.ID),
		Title: result.Title,
	}

	return response, nil
}
