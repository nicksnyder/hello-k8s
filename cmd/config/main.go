package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/nicksnyder/hello-server/pkg/debug"
	"github.com/nicksnyder/hello-server/pkg/env"
	apphttp "github.com/nicksnyder/hello-server/pkg/http"
)

var configFile string
var configMap string

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

	configMap = os.Getenv("CONFIG_MAP")
	if path == "" {
		panic("CONFIG_MAP required")
	}

	server := apphttp.NewServeMux()
	server.HandleErrFunc("/file", handleFile)
	server.HandleErrFunc("/configmap", handleConfigMap)

	log.Printf("listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatal(err)
	}
}

func handleConfigMap(w http.ResponseWriter, r *http.Request) error {
	b, err := ioutil.ReadFile(configMap)
	if err != nil {
		return err
	}
	if err := debug.WriteData(w); err != nil {
		return err
	}
	w.Write([]byte("configmap"))
	_, err = w.Write(b)
	return err
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
