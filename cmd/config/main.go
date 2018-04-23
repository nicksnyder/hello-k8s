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
	port := env.MustGet("CONFIG_PORT")

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
	server.HandleErrFunc("/file", handleFile)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleFile(w http.ResponseWriter, r *http.Request) error {
	if write := r.FormValue("write"); write != "" {
		if err := ioutil.WriteFile(configFile, []byte(write), 0644); err != nil {
			return err
		}
	}
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	if err := debug.WriteData(w); err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
