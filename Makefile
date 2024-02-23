.PHONY: dev build-image migrate

dev:
	dotenv wgo run ./cmd/server

build:
	go build -ldflags '-s -w' -o ./bin/server ./cmd/server

build-image:
	docker buildx build --platform=linux/amd64 -t at-feeds .

migrate:
	migrate -database ${DATABASE_URL} -path migrations/sqlite up

lint:
	golangci-lint run

format:
	gofumpt -l -w .