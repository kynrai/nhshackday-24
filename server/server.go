package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	r       *chi.Mux
	tw      *Twilio
	ian     *IanClient
	msgChan chan string
}

func NewServer(conf Config) (*Server, error) {
	return &Server{
		r:       chi.NewRouter(),
		tw:      NewTwilio(conf.TwilioSID, conf.TwilioAuth, conf.TwilioFrom, conf.TwilioTo),
		ian:     NewIanClient(),
		msgChan: make(chan string),
	}, nil
}

func (s *Server) Start() {
	s.Routes()
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", s.r)
}
