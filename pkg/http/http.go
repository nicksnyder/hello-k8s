package http

import (
	"log"
	"net/http"

	"github.com/nicksnyder/service/pkg/debug"
)

type ServeMux struct {
	http.ServeMux
}

func NewServeMux() *ServeMux {
	m := &ServeMux{}
	m.HandleFunc("/healthz", handleHealthz)
	m.HandleErrFunc("/debug", debug.Serve)
	return m
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
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
