FROM node:alpine AS react-build
ENV VITE_BACKEND_URL=/api
WORKDIR /app
COPY web/package*.json ./
RUN npm ci
COPY web ./
RUN npm run build

FROM golang:1.22-alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
COPY --from=react-build /app/dist ./web/dist
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' ./cmd/server

FROM gcr.io/distroless/static AS release
COPY --from=go-build /app/server /usr/local/bin/at-feeds
ENTRYPOINT ["at-feeds"]