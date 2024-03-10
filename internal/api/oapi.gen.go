// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Collection defines model for Collection.
type Collection struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// CreateCollection defines model for CreateCollection.
type CreateCollection struct {
	ParentId *int   `json:"parentId,omitempty"`
	Title    string `json:"title"`
}

// CreateFeed defines model for CreateFeed.
type CreateFeed struct {
	Url string `json:"url"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Feed defines model for Feed.
type Feed struct {
	EntryCount  *int    `json:"entryCount,omitempty"`
	Id          int     `json:"id"`
	Link        string  `json:"link"`
	Title       string  `json:"title"`
	UnreadCount int     `json:"unreadCount"`
	Url         *string `json:"url"`
}

// FeedEntry defines model for FeedEntry.
type FeedEntry struct {
	Author       *string   `json:"author"`
	Content      *string   `json:"content"`
	FeedId       int       `json:"feedId"`
	HasRead      bool      `json:"hasRead"`
	Id           int       `json:"id"`
	Link         string    `json:"link"`
	PublishedAt  time.Time `json:"publishedAt"`
	ThumbnailUrl *string   `json:"thumbnailUrl"`
	Title        string    `json:"title"`
}

// File defines model for File.
type File = openapi_types.File

// UpdateFeedEntry defines model for UpdateFeedEntry.
type UpdateFeedEntry struct {
	HasRead *bool `json:"hasRead,omitempty"`
}

// ListCollectionsParams defines parameters for ListCollections.
type ListCollectionsParams struct {
	Limit    *int `form:"limit,omitempty" json:"limit,omitempty"`
	Page     *int `form:"page,omitempty" json:"page,omitempty"`
	ParentId *int `form:"parentId,omitempty" json:"parentId,omitempty"`
}

// ListFeedEntriesParams defines parameters for ListFeedEntries.
type ListFeedEntriesParams struct {
	FeedId  *int  `form:"feedId,omitempty" json:"feedId,omitempty"`
	HasRead *bool `form:"hasRead,omitempty" json:"hasRead,omitempty"`
	Limit   *int  `form:"limit,omitempty" json:"limit,omitempty"`
	Page    *int  `form:"page,omitempty" json:"page,omitempty"`
}

// ListFeedsParams defines parameters for ListFeeds.
type ListFeedsParams struct {
	Limit        *int `form:"limit,omitempty" json:"limit,omitempty"`
	Page         *int `form:"page,omitempty" json:"page,omitempty"`
	CollectionId *int `form:"collectionId,omitempty" json:"collectionId,omitempty"`
}

// CreateCollectionJSONRequestBody defines body for CreateCollection for application/json ContentType.
type CreateCollectionJSONRequestBody = CreateCollection

// CreateFeedJSONRequestBody defines body for CreateFeed for application/json ContentType.
type CreateFeedJSONRequestBody = CreateFeed

// UpdateFeedEntryJSONRequestBody defines body for UpdateFeedEntry for application/json ContentType.
type UpdateFeedEntryJSONRequestBody = UpdateFeedEntry

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /collections)
	ListCollections(w http.ResponseWriter, r *http.Request, params ListCollectionsParams)

	// (POST /collections)
	CreateCollection(w http.ResponseWriter, r *http.Request)

	// (GET /entries)
	ListFeedEntries(w http.ResponseWriter, r *http.Request, params ListFeedEntriesParams)

	// (GET /feeds)
	ListFeeds(w http.ResponseWriter, r *http.Request, params ListFeedsParams)

	// (POST /feeds)
	CreateFeed(w http.ResponseWriter, r *http.Request)

	// (POST /feeds/export)
	ExportFeeds(w http.ResponseWriter, r *http.Request)

	// (POST /feeds/import)
	ImportFeeds(w http.ResponseWriter, r *http.Request)

	// (PATCH /feeds/{feedId}/entries/{entryId})
	UpdateFeedEntry(w http.ResponseWriter, r *http.Request, feedId int, entryId int)

	// (DELETE /feeds/{id})
	DeleteFeed(w http.ResponseWriter, r *http.Request, id int)

	// (GET /feeds/{id})
	GetFeed(w http.ResponseWriter, r *http.Request, id int)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (GET /collections)
func (_ Unimplemented) ListCollections(w http.ResponseWriter, r *http.Request, params ListCollectionsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /collections)
func (_ Unimplemented) CreateCollection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /entries)
func (_ Unimplemented) ListFeedEntries(w http.ResponseWriter, r *http.Request, params ListFeedEntriesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /feeds)
func (_ Unimplemented) ListFeeds(w http.ResponseWriter, r *http.Request, params ListFeedsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /feeds)
func (_ Unimplemented) CreateFeed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /feeds/export)
func (_ Unimplemented) ExportFeeds(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /feeds/import)
func (_ Unimplemented) ImportFeeds(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (PATCH /feeds/{feedId}/entries/{entryId})
func (_ Unimplemented) UpdateFeedEntry(w http.ResponseWriter, r *http.Request, feedId int, entryId int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (DELETE /feeds/{id})
func (_ Unimplemented) DeleteFeed(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (GET /feeds/{id})
func (_ Unimplemented) GetFeed(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListCollections operation middleware
func (siw *ServerInterfaceWrapper) ListCollections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListCollectionsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "parentId" -------------

	err = runtime.BindQueryParameter("form", true, false, "parentId", r.URL.Query(), &params.ParentId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "parentId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListCollections(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateCollection operation middleware
func (siw *ServerInterfaceWrapper) CreateCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateCollection(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListFeedEntries operation middleware
func (siw *ServerInterfaceWrapper) ListFeedEntries(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListFeedEntriesParams

	// ------------- Optional query parameter "feedId" -------------

	err = runtime.BindQueryParameter("form", true, false, "feedId", r.URL.Query(), &params.FeedId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feedId", Err: err})
		return
	}

	// ------------- Optional query parameter "hasRead" -------------

	err = runtime.BindQueryParameter("form", true, false, "hasRead", r.URL.Query(), &params.HasRead)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "hasRead", Err: err})
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListFeedEntries(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListFeeds operation middleware
func (siw *ServerInterfaceWrapper) ListFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListFeedsParams

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "collectionId" -------------

	err = runtime.BindQueryParameter("form", true, false, "collectionId", r.URL.Query(), &params.CollectionId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "collectionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListFeeds(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateFeed operation middleware
func (siw *ServerInterfaceWrapper) CreateFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateFeed(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ExportFeeds operation middleware
func (siw *ServerInterfaceWrapper) ExportFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ExportFeeds(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ImportFeeds operation middleware
func (siw *ServerInterfaceWrapper) ImportFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportFeeds(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateFeedEntry operation middleware
func (siw *ServerInterfaceWrapper) UpdateFeedEntry(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "feedId" -------------
	var feedId int

	err = runtime.BindStyledParameterWithOptions("simple", "feedId", chi.URLParam(r, "feedId"), &feedId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "feedId", Err: err})
		return
	}

	// ------------- Path parameter "entryId" -------------
	var entryId int

	err = runtime.BindStyledParameterWithOptions("simple", "entryId", chi.URLParam(r, "entryId"), &entryId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "entryId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateFeedEntry(w, r, feedId, entryId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteFeed operation middleware
func (siw *ServerInterfaceWrapper) DeleteFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteFeed(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetFeed operation middleware
func (siw *ServerInterfaceWrapper) GetFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFeed(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/collections", wrapper.ListCollections)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/collections", wrapper.CreateCollection)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/entries", wrapper.ListFeedEntries)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/feeds", wrapper.ListFeeds)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/feeds", wrapper.CreateFeed)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/feeds/export", wrapper.ExportFeeds)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/feeds/import", wrapper.ImportFeeds)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/feeds/{feedId}/entries/{entryId}", wrapper.UpdateFeedEntry)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/feeds/{id}", wrapper.DeleteFeed)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/feeds/{id}", wrapper.GetFeed)
	})

	return r
}

type ListCollectionsRequestObject struct {
	Params ListCollectionsParams
}

type ListCollectionsResponseObject interface {
	VisitListCollectionsResponse(w http.ResponseWriter) error
}

type ListCollections200JSONResponse struct {
	Data    []Collection `json:"data"`
	HasMore bool         `json:"hasMore"`
}

func (response ListCollections200JSONResponse) VisitListCollectionsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateCollectionRequestObject struct {
	Body *CreateCollectionJSONRequestBody
}

type CreateCollectionResponseObject interface {
	VisitCreateCollectionResponse(w http.ResponseWriter) error
}

type CreateCollection201JSONResponse Collection

func (response CreateCollection201JSONResponse) VisitCreateCollectionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateCollection400JSONResponse Error

func (response CreateCollection400JSONResponse) VisitCreateCollectionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ListFeedEntriesRequestObject struct {
	Params ListFeedEntriesParams
}

type ListFeedEntriesResponseObject interface {
	VisitListFeedEntriesResponse(w http.ResponseWriter) error
}

type ListFeedEntries200JSONResponse struct {
	Data    []FeedEntry `json:"data"`
	HasMore bool        `json:"hasMore"`
}

func (response ListFeedEntries200JSONResponse) VisitListFeedEntriesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type ListFeedsRequestObject struct {
	Params ListFeedsParams
}

type ListFeedsResponseObject interface {
	VisitListFeedsResponse(w http.ResponseWriter) error
}

type ListFeeds200JSONResponse struct {
	Data    []Feed `json:"data"`
	HasMore bool   `json:"hasMore"`
}

func (response ListFeeds200JSONResponse) VisitListFeedsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateFeedRequestObject struct {
	Body *CreateFeedJSONRequestBody
}

type CreateFeedResponseObject interface {
	VisitCreateFeedResponse(w http.ResponseWriter) error
}

type CreateFeed201JSONResponse Feed

func (response CreateFeed201JSONResponse) VisitCreateFeedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateFeed400JSONResponse Error

func (response CreateFeed400JSONResponse) VisitCreateFeedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ExportFeedsRequestObject struct {
}

type ExportFeedsResponseObject interface {
	VisitExportFeedsResponse(w http.ResponseWriter) error
}

type ExportFeeds200ApplicationoctetStreamResponse struct {
	Body          io.Reader
	ContentLength int64
}

func (response ExportFeeds200ApplicationoctetStreamResponse) VisitExportFeedsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	if response.ContentLength != 0 {
		w.Header().Set("Content-Length", fmt.Sprint(response.ContentLength))
	}
	w.WriteHeader(200)

	if closer, ok := response.Body.(io.ReadCloser); ok {
		defer closer.Close()
	}
	_, err := io.Copy(w, response.Body)
	return err
}

type ExportFeeds500JSONResponse Error

func (response ExportFeeds500JSONResponse) VisitExportFeedsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ImportFeedsRequestObject struct {
	Body io.Reader
}

type ImportFeedsResponseObject interface {
	VisitImportFeedsResponse(w http.ResponseWriter) error
}

type ImportFeeds200Response struct {
}

func (response ImportFeeds200Response) VisitImportFeedsResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type ImportFeeds500JSONResponse Error

func (response ImportFeeds500JSONResponse) VisitImportFeedsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type UpdateFeedEntryRequestObject struct {
	FeedId  int `json:"feedId"`
	EntryId int `json:"entryId"`
	Body    *UpdateFeedEntryJSONRequestBody
}

type UpdateFeedEntryResponseObject interface {
	VisitUpdateFeedEntryResponse(w http.ResponseWriter) error
}

type UpdateFeedEntry200JSONResponse FeedEntry

func (response UpdateFeedEntry200JSONResponse) VisitUpdateFeedEntryResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateFeedEntry404JSONResponse Error

func (response UpdateFeedEntry404JSONResponse) VisitUpdateFeedEntryResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type DeleteFeedRequestObject struct {
	Id int `json:"id"`
}

type DeleteFeedResponseObject interface {
	VisitDeleteFeedResponse(w http.ResponseWriter) error
}

type DeleteFeed204Response struct {
}

func (response DeleteFeed204Response) VisitDeleteFeedResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteFeed404JSONResponse Error

func (response DeleteFeed404JSONResponse) VisitDeleteFeedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetFeedRequestObject struct {
	Id int `json:"id"`
}

type GetFeedResponseObject interface {
	VisitGetFeedResponse(w http.ResponseWriter) error
}

type GetFeed200JSONResponse Feed

func (response GetFeed200JSONResponse) VisitGetFeedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetFeed404JSONResponse Error

func (response GetFeed404JSONResponse) VisitGetFeedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /collections)
	ListCollections(ctx context.Context, request ListCollectionsRequestObject) (ListCollectionsResponseObject, error)

	// (POST /collections)
	CreateCollection(ctx context.Context, request CreateCollectionRequestObject) (CreateCollectionResponseObject, error)

	// (GET /entries)
	ListFeedEntries(ctx context.Context, request ListFeedEntriesRequestObject) (ListFeedEntriesResponseObject, error)

	// (GET /feeds)
	ListFeeds(ctx context.Context, request ListFeedsRequestObject) (ListFeedsResponseObject, error)

	// (POST /feeds)
	CreateFeed(ctx context.Context, request CreateFeedRequestObject) (CreateFeedResponseObject, error)

	// (POST /feeds/export)
	ExportFeeds(ctx context.Context, request ExportFeedsRequestObject) (ExportFeedsResponseObject, error)

	// (POST /feeds/import)
	ImportFeeds(ctx context.Context, request ImportFeedsRequestObject) (ImportFeedsResponseObject, error)

	// (PATCH /feeds/{feedId}/entries/{entryId})
	UpdateFeedEntry(ctx context.Context, request UpdateFeedEntryRequestObject) (UpdateFeedEntryResponseObject, error)

	// (DELETE /feeds/{id})
	DeleteFeed(ctx context.Context, request DeleteFeedRequestObject) (DeleteFeedResponseObject, error)

	// (GET /feeds/{id})
	GetFeed(ctx context.Context, request GetFeedRequestObject) (GetFeedResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// ListCollections operation middleware
func (sh *strictHandler) ListCollections(w http.ResponseWriter, r *http.Request, params ListCollectionsParams) {
	var request ListCollectionsRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListCollections(ctx, request.(ListCollectionsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListCollections")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListCollectionsResponseObject); ok {
		if err := validResponse.VisitListCollectionsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateCollection operation middleware
func (sh *strictHandler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	var request CreateCollectionRequestObject

	var body CreateCollectionJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateCollection(ctx, request.(CreateCollectionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateCollection")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateCollectionResponseObject); ok {
		if err := validResponse.VisitCreateCollectionResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListFeedEntries operation middleware
func (sh *strictHandler) ListFeedEntries(w http.ResponseWriter, r *http.Request, params ListFeedEntriesParams) {
	var request ListFeedEntriesRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListFeedEntries(ctx, request.(ListFeedEntriesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListFeedEntries")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListFeedEntriesResponseObject); ok {
		if err := validResponse.VisitListFeedEntriesResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ListFeeds operation middleware
func (sh *strictHandler) ListFeeds(w http.ResponseWriter, r *http.Request, params ListFeedsParams) {
	var request ListFeedsRequestObject

	request.Params = params

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ListFeeds(ctx, request.(ListFeedsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ListFeeds")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ListFeedsResponseObject); ok {
		if err := validResponse.VisitListFeedsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateFeed operation middleware
func (sh *strictHandler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	var request CreateFeedRequestObject

	var body CreateFeedJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateFeed(ctx, request.(CreateFeedRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateFeed")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateFeedResponseObject); ok {
		if err := validResponse.VisitCreateFeedResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ExportFeeds operation middleware
func (sh *strictHandler) ExportFeeds(w http.ResponseWriter, r *http.Request) {
	var request ExportFeedsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ExportFeeds(ctx, request.(ExportFeedsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ExportFeeds")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ExportFeedsResponseObject); ok {
		if err := validResponse.VisitExportFeedsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// ImportFeeds operation middleware
func (sh *strictHandler) ImportFeeds(w http.ResponseWriter, r *http.Request) {
	var request ImportFeedsRequestObject

	request.Body = r.Body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ImportFeeds(ctx, request.(ImportFeedsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ImportFeeds")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ImportFeedsResponseObject); ok {
		if err := validResponse.VisitImportFeedsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateFeedEntry operation middleware
func (sh *strictHandler) UpdateFeedEntry(w http.ResponseWriter, r *http.Request, feedId int, entryId int) {
	var request UpdateFeedEntryRequestObject

	request.FeedId = feedId
	request.EntryId = entryId

	var body UpdateFeedEntryJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateFeedEntry(ctx, request.(UpdateFeedEntryRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateFeedEntry")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(UpdateFeedEntryResponseObject); ok {
		if err := validResponse.VisitUpdateFeedEntryResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteFeed operation middleware
func (sh *strictHandler) DeleteFeed(w http.ResponseWriter, r *http.Request, id int) {
	var request DeleteFeedRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteFeed(ctx, request.(DeleteFeedRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteFeed")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteFeedResponseObject); ok {
		if err := validResponse.VisitDeleteFeedResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetFeed operation middleware
func (sh *strictHandler) GetFeed(w http.ResponseWriter, r *http.Request, id int) {
	var request GetFeedRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFeed(ctx, request.(GetFeedRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFeed")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFeedResponseObject); ok {
		if err := validResponse.VisitGetFeedResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
