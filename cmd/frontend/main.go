package main

import (
	"io"
	"log"
	"net/http"

	"github.com/nicksnyder/service/pkg/debug"
	"github.com/nicksnyder/service/pkg/env"
)

func main() {
	port := env.Get("FRONTEND_PORT", "8080")

	server := http.NewServeMux()
	server.HandleFunc("/debug", debug.Serve)
	server.HandleFunc("/config", handleConfig)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleConfig(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://config/file")
	if err != nil {
		log.Println("couldn't fetch config", err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Println("failed to copy config", err)
	}
}
