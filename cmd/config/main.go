package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/nicksnyder/service/pkg/debug"
	"github.com/nicksnyder/service/pkg/env"
	apphttp "github.com/nicksnyder/service/pkg/http"
)

var configFile string

func main() {
	port := env.Get("CONFIG_PORT", "7070")

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		u, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = u.HomeDir
	}
	configFile = filepath.Join(path, "config.json")

	server := apphttp.NewServeMux()
	server.HandleErrFunc("/debug", debug.Serve)
	server.HandleErrFunc("/file", handleFile)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleFile(w http.ResponseWriter, r *http.Request) error {
	if write := r.FormValue("write"); write != "" {
		return ioutil.WriteFile(configFile, []byte(write), 0644)
	}
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
