package server

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/kynrai/nhshackday-24/ui/patient"
)

func (s *Server) handleNofityConnect(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	s.notifyMsg = make(chan string)

	// Send data to the client
	go func() {
		for data := range s.notifyMsg {
			fmt.Fprintf(w, "data: %s\n\n", data)
			if f, ok := w.(http.Flusher); ok {
				if f != nil {
					f.Flush()
				}
			}
		}
	}()
	<-r.Context().Done()
	fmt.Printf("Client disconnected: %v\n", r.Context().Err())
}

func (s *Server) handleSendAlert(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	patient.Alert().Render(r.Context(), buf)
	s.notifyMsg <- buf.String()
}

func (s *Server) handleSendAppointment(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	patient.Appointment().Render(r.Context(), buf)
	s.notifyMsg <- buf.String()
}
