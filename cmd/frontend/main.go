package main

import (
	"io"
	"log"
	"net/http"

	"github.com/nicksnyder/service/pkg/debug"
	"github.com/nicksnyder/service/pkg/env"
	apphttp "github.com/nicksnyder/service/pkg/http"
)

func main() {
	port := env.Get("FRONTEND_PORT", "8080")

	server := apphttp.NewServeMux()
	server.HandleErrFunc("/debug", debug.Serve)
	server.HandleErrFunc("/config", handleConfig)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleConfig(w http.ResponseWriter, r *http.Request) error {
	resp, err := http.Get("http://config/file")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}
	return nil
}
