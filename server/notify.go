package server

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/kynrai/nhshackday-24/ui"
	"github.com/kynrai/nhshackday-24/ui/patient"
)

func (s *Server) handlePageSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.SSE()).Render(r.Context(), w)
}

func (s *Server) handlePageSSESend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.SSESend()).Render(r.Context(), w)
}

func (s *Server) handleSSEConnect(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	s.msgChan = make(chan string)

	// Send data to the client
	go func() {
		for data := range s.msgChan {
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

	// Simulate sending data periodically
	// for {
	// 	s.msgChan <- time.Now().Format(time.TimeOnly)
	// 	time.Sleep(1 * time.Second)
	// }
}

func (s *Server) handleNotify(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	patient.Alert().Render(r.Context(), buf)
	s.msgChan <- buf.String()
	// s.msgChan <- fmt.Sprintf("data: %s\n\n", buf.String())
}
