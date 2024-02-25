package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/kynrai/nhshackday-24/model"
)

type Server struct {
	r         *chi.Mux
	tw        *Twilio
	ian       *IanClient
	notifyMsg chan string
	reports   []model.SingleMtxReport
}

func NewServer(conf Config) (*Server, error) {
	return &Server{
		r:         chi.NewRouter(),
		tw:        NewTwilio(conf.TwilioSID, conf.TwilioAuth, conf.TwilioFrom, conf.TwilioTo),
		ian:       NewIanClient(),
		notifyMsg: make(chan string),
		reports:   model.MtxReports(),
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
