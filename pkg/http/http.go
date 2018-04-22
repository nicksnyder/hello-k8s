package http

import (
	"log"
	"net/http"
)

type ServeMux struct {
	http.ServeMux
}

func NewServeMux() *ServeMux {
	return &ServeMux{}
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
