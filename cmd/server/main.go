package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/abhithube/at-feeds/internal/api"
	"github.com/abhithube/at-feeds/internal/backup"
	"github.com/abhithube/at-feeds/internal/database"
	"github.com/abhithube/at-feeds/internal/handler"
	"github.com/abhithube/at-feeds/internal/task"
	"github.com/abhithube/at-feeds/migrations"
	"github.com/abhithube/at-feeds/plugins"
	"github.com/hashicorp/go-retryablehttp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron/v3"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	port := 8000
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}
	shouldRefresh := true
	if disableRefreshStr := os.Getenv("DISABLE_REFRESH"); disableRefreshStr != "" {
		if d, err := strconv.ParseBool(disableRefreshStr); err == nil {
			shouldRefresh = !d
		}
	}
	frontendURL := os.Getenv("FRONTEND_URL")

	db, err := sql.Open("sqlite3", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = migrations.Migrate(db); err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)

	retryClient := retryablehttp.NewClient()
	httpClient := retryClient.StandardClient()

	manager := task.NewManager(db, queries, httpClient)
	worker := task.NewWorker(manager)
	plugins.Register()

	backupManager := backup.NewOPMLManager()

	router := http.NewServeMux()

	si := api.NewStrictHandler(handler.New(db, queries, worker, backupManager), nil)

	handler := api.HandlerFromMuxWithBaseURL(si, router, "/api")

	if frontendURL == "" {
		router.Handle("GET /", api.SPAHandler("dist"))
	}

	handler = api.CORSHandler(frontendURL)(handler)

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	if shouldRefresh {
		log.Println("Initializing refresh job...")

		job := task.NewRefreshJob(queries, worker)
		c := cron.New()
		if _, err = c.AddFunc("*/15 * * * *", func() { job.Run(context.Background()) }); err != nil {
			log.Fatal(err)
		}

		c.Start()
	}

	log.Printf("Starting server at address %s\n", httpServer.Addr)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
