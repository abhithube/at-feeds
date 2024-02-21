package internal

//go:generate sqlc generate -f ../sqlc.yml
//go:generate oapi-codegen -generate models,chi-server,strict-server -o ./api/oapi.gen.go -package api ../openapi.json
