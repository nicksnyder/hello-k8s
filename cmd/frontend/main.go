package main

import (
	"io"
	"log"
	"net/http"

	"github.com/nicksnyder/service/pkg/env"
	apphttp "github.com/nicksnyder/service/pkg/http"
)

func main() {
	port := env.MustGet("FRONTEND_PORT")

	server := apphttp.NewServeMux()
	server.HandleErrFunc("/config", handleConfig)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleConfig(w http.ResponseWriter, r *http.Request) error {
	resp, err := http.Get("http://config/file?write=" + r.FormValue("write"))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		return err
	}
	return nil
}
