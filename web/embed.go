package web

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/*
var BuildFs embed.FS

func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(BuildFs, "dist")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(build)
}
