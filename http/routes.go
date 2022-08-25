package http

import (
	"fmt"
	"net/http"

	"github.com/madxiii/axxonsofttask/service"
)

var wrongMethod []byte = []byte("only GET method is allowed")

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Index)
	mux.HandleFunc("/response", h.Response)
	mux.HandleFunc("/request", h.Request)

	return mux
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Printf("err: %s\n", message)
	w.Write([]byte(message))
}

func okResponse(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	w.Write(body)
}

func wrongMethodResponse(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	w.Write(body)
}
