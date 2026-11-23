package http

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	addr string
	mux  *http.ServeMux
}

func New(addr string, mux *http.ServeMux) *Server{
	return &Server{
		addr: addr,
		mux	:  mux,
	}
}

func (s*Server) Start() error{
	httpServer := &http.Server{
		Addr		: s.addr,
		Handler		: s.mux,
		ReadTimeout	: 10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	log.Printf("Server jalan di http://localhost%s\n", s.addr)
	return httpServer.ListenAndServe()
}