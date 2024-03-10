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
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	parentID := request.Params.ParentId
	limit := request.Params.Limit
	page := request.Params.Page

	params := database.ListCollectionsParams{Limit: -1}
	if limit != nil {
		params.Limit = int64(*limit)
		if page != nil {
			params.Offset = (int64(*page) - 1) * params.Limit
		}
	}
	if parentID != nil {
		params.FilterByParentID = true
		params.ParentID = parentID
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
	parentID := request.Body.ParentId

	if len(title) == 0 {
		return api.CreateCollection400JSONResponse{Message: "'title' cannot be empty"}, nil
	}

	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer database.Rollback(tx)

	qtx := h.queries.WithTx(tx)

	params := database.InsertCollectionParams{
		Title: title,
	}
	if parentID != nil {
		params.ParentID = sql.NullInt64{Int64: int64(*parentID), Valid: true}

		_, err := qtx.GetCollection(ctx, int64(*parentID))
		if err != nil {
			return api.CreateCollection400JSONResponse{Message: "Invalid parent ID"}, nil
		}
	}
	result, err := qtx.InsertCollection(ctx, params)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, sql.ErrNoRows) {
			msg = fmt.Sprintf("Collection already exists with title '%s' and parent '%d'", params.Title, params.ParentID.Int64)
		}

		return api.CreateCollection400JSONResponse{Message: msg}, nil
	}

	response := api.CreateCollection201JSONResponse{
		Id:    int(result.ID),
		Title: result.Title,
	}

	return response, tx.Commit()
}
