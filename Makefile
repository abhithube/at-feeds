.PHONY: dev build-image migrate

dev:
	dotenv wgo run ./cmd/server

build:
	go build -o ./bin/server -ldflags '-s -w' ./cmd/server

build-image:
	docker buildx build --platform=linux/amd64 -t at-feeds .

migrate:
	migrate -path migrations -database 'sqlite3://dev.db' up

format:
	gofumpt -l -w .