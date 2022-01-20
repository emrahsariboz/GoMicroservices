package handlers

import (
	"log"
	"net/http"
)

type GoodByte struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Handler {
	return &Handler{l}
}

func (h *GoodByte) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("GoodBye"))
}
