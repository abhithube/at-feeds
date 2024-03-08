package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/abhithube/at-feeds/web"
)

func SPAHandler(buildPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := web.BuildFs.Open(filepath.Join(buildPath, r.URL.Path))
		if os.IsNotExist(err) {
			index, err := web.BuildFs.ReadFile(filepath.Join(buildPath, "index.html"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusAccepted)

			if _, err = w.Write(index); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		defer f.Close()

		http.FileServer(web.BuildHTTPFS()).ServeHTTP(w, r)
	})
}
