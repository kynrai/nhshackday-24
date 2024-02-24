package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
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
	s.r.Get("/ws", s.handleWS)
	s.r.Get("/sms", s.handleSendSMS)
	s.r.Get("/dummy", s.handleDummy)
	s.r.Get("/clinician-view", s.handleClinicianView)
}
