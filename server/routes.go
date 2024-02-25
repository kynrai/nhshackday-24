package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed assets
var assets embed.FS

func (s *Server) Routes() {
	contentAssets, err := fs.Sub(fs.FS(assets), "assets")
	if err != nil {
		log.Fatal(err)
	}
	s.r.Method(http.MethodGet, "/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(contentAssets))))
	s.r.Get("/", s.handlePing)
	s.r.Post("/", s.handleSubmit)
	s.r.Get("/sms", s.handleSendSMS)
	s.r.Get("/dummy", s.handleDummy)
	s.r.Get("/clinician", s.handleClinicianView)
	s.r.Get("/patient", s.handlePatientView)
	s.r.Route("/hx", func(r chi.Router) {
		r.Get("/clinician-nav/{type}", s.handleClinicianNav)
		r.Get("/sms-reminder", s.handleSMSReminder)
		r.Get("/sms-alert", s.handleSMSAlert)
		r.Get("/patient-nav/{type}", s.handlePatientNav)
	})

	s.r.HandleFunc("/notify", s.handleNofityConnect)
	s.r.Post("/app-alert", s.handleSendAlert)
	s.r.Post("/app-reminder", s.handleSendAppointment)
}
