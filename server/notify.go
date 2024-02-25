package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kynrai/nhshackday-24/ui"
)

func (s *Server) handlePageSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.SSE()).Render(r.Context(), w)
}

func (s *Server) handleSSEConnect(w http.ResponseWriter, r *http.Request) {
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
}

func (s *Server) handleNotify(w http.ResponseWriter, r *http.Request) {

}
