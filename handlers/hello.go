package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello World!")
	d, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Ooops"))
		return
	}
	// Response Writer is an interface, used to write a response to request!
	fmt.Fprintf(rw, "GoodBye %s", d)
}
