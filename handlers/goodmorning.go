package handlers

import (
	"log"
	"net/http"
)

type MorningHandler struct{}

func NewMorningHanlder() MorningHandler {
	return MorningHandler{}
}

func (m MorningHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("serveHTTP has been called")

	rw.Write([]byte("This is just test. I am writing to response."))

}
