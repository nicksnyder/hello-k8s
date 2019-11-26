package http

import (
	"log"
	"net/http"

	"github.com/nicksnyder/hello-server/cmd/internal/debug"
)

type ServeMux struct {
	http.ServeMux
}

func NewServeMux() *ServeMux {
	m := &ServeMux{}
	m.HandleErrFunc("/", handleHello)
	return m
}

func handleHello(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Hello from\n"))
	if err != nil {
		return err
	}
	return debug.WriteData(w)
}

func (mux *ServeMux) HandleErrFunc(pattern string, handlerErrFunc HandlerErrFunc) {
	mux.Handle(pattern, handlerErrFunc)
}

// HandlerErrFunc handles a HTTP request.
type HandlerErrFunc func(w http.ResponseWriter, r *http.Request) error

func (h HandlerErrFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
	}
}
