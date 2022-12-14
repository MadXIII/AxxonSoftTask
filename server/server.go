package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(address string, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        handlers,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	fmt.Printf("Server is listening%s\n", address)
	return s.httpServer.ListenAndServe()
}
