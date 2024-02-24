package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	r http.ServeMux
}

func NewServer() *Server {
	return &Server{
		r: *http.NewServeMux(),
	}
}

func (s *Server) Start() {
	s.Routes()
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", &s.r)
}
