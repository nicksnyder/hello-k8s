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

	server := http.NewServeMux()
	server.HandleFunc("/debug", debug.Serve)
	http.HandleFunc("/file", handleFile)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}
