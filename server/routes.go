package server

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

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
	s.r.Get("/ws", s.handleWS)
	s.r.Get("/sms", s.handleSendSMS)
	s.r.Get("/dummy", s.handleDummy)
	s.r.Post("/notify", s.handleNotify)
	s.r.Get("/sse", s.handlePageSSE)
	s.r.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// Set headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create a channel to send data
		dataCh := make(chan string)

		// Create a context for handling client disconnection
		_, cancel := context.WithCancel(r.Context())
		defer cancel()

		// Send data to the client
		go func() {
			for data := range dataCh {
				fmt.Fprintf(w, "data: %s\n\n", data)
				w.(http.Flusher).Flush()
			}
		}()

		// Simulate sending data periodically
		for {
			dataCh <- time.Now().Format(time.TimeOnly)
			time.Sleep(1 * time.Second)
		}
	})

	s.r.Get("/notifications", s.handlePageNotifications)
	s.r.Get("/cons", s.handlePageConnections)

	s.r.Get("/clinician", s.handleClinicianView)
	s.r.Route("/hx", func(r chi.Router) {
		r.Get("/nav/{type}", s.handleNav)
	})
}
