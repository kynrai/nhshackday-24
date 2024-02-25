package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), s.r))
}
