package api

import (
	"net/http"

	"github.com/rs/cors"
)

func CORSHandler(origin string) func(next http.Handler) http.Handler {
	options := cors.Options{
		AllowedOrigins: []string{origin},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
	}

	return cors.New(options).Handler
}
