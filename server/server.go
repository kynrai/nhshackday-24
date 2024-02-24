package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Server struct {
	r       *chi.Mux
	tw      *Twilio
	ian     *IanClient
	clients map[string]*websocket.Conn
}

func NewServer(conf Config) (*Server, error) {
	return &Server{
		r:       chi.NewRouter(),
		tw:      NewTwilio(conf.TwilioSID, conf.TwilioAuth, conf.TwilioFrom, conf.TwilioTo),
		ian:     NewIanClient(),
		clients: make(map[string]*websocket.Conn),
	}, nil
}

func (s *Server) Start() {
	s.Routes()
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", s.r)
}
