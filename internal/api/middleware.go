package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/abhithube/at-feeds/web"
	"github.com/go-chi/cors"
)

func CorsHandler(origin string) func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{origin},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
	})
}

func HandleSPA(buildPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := web.BuildFs.Open(filepath.Join(buildPath, r.URL.Path))
		if os.IsNotExist(err) {
			index, err := web.BuildFs.ReadFile(filepath.Join(buildPath, "index.html"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusAccepted)
			w.Write(index)

			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		defer f.Close()

		http.FileServer(web.BuildHTTPFS()).ServeHTTP(w, r)
	})
}
