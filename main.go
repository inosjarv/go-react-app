package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed client/dist/*
var embeddedFiles embed.FS

func main() {
	distDir, err := fs.Sub(embeddedFiles, "client/dist")
	if err != nil {
		log.Fatalf("Failed to locate embedded dist folder: %v", err)
	}

	http.Handle("/", http.FileServer(http.FS(distDir)))

	port := ":8080"
	log.Printf("Serving Vite assets on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
