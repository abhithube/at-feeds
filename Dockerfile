FROM node:alpine AS react-build
ENV VITE_BACKEND_URL=/api
WORKDIR /app
COPY web/package*.json ./
RUN npm ci
COPY web ./
RUN npm run build

FROM golang:1.22-alpine AS go-build
RUN apk add --no-cache --update go gcc g++
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
COPY --from=react-build /app/dist ./web/dist
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags '-s -w' ./cmd/server

FROM alpine AS release
RUN apk add --no-cache --update sqlite
COPY --from=go-build /app/server /usr/local/bin/at-feeds
VOLUME ["/data"]
ENV DATABASE_URL=/data/db.sqlite3
ENTRYPOINT ["at-feeds"]