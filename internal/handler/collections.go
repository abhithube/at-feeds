package handler

import (
	"context"
	"database/sql"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/database"
)

func (h *Handler) ListCollections(ctx context.Context, request api.ListCollectionsRequestObject) (api.ListCollectionsResponseObject, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	params := database.ListCollectionsParams{Limit: -1}
	if request.Params.Limit != nil {
		params.Limit = int64(*request.Params.Limit)
		if request.Params.Page != nil {
			params.Offset = (int64(*request.Params.Page) - 1) * params.Limit
		}
	}
	if request.Params.ParentId != nil {
		params.FilterByParentID = true
		params.ParentID = sql.NullInt64{Int64: int64(*request.Params.ParentId), Valid: true}
	}

	result, err := qtx.ListCollections(ctx, params)
	if err != nil {
		return nil, err
	}

	params2 := database.CountCollectionsParams{
		FilterByParentID: params.FilterByParentID,
		ParentID:         params.ParentID,
	}
	count, err := qtx.CountCollections(ctx, params2)
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

	response := api.ListCollections200JSONResponse{
		Data:    arr,
		HasMore: (params.Limit + params.Offset) < count,
	}

	return response, tx.Commit()
}

func (h *Handler) CreateCollection(ctx context.Context, request api.CreateCollectionRequestObject) (api.CreateCollectionResponseObject, error) {
	title := request.Body.Title
	if len(title) == 0 {
		return api.CreateCollection400JSONResponse{Message: "'title' cannot be empty"}, nil
	}

	params := database.InsertCollectionParams{
		Title: title,
	}
	result, err := h.queries.InsertCollection(ctx, params)
	if err != nil {
		return api.CreateCollection400JSONResponse{Message: err.Error()}, nil
	}

	response := api.CreateCollection201JSONResponse{
		Id:    int(result.ID),
		Title: result.Title,
	}

	return response, nil
}
