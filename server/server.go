package server

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()

	return &Server{
		mux: mux,
	}
}

func (s *Server) Init() error {
	addRoutes(s.mux)
	return nil
}

func (s *Server) Mux() http.ServeMux {
	return *s.mux
}

func (s *Server) Shutdown() []error {
	return nil
}
