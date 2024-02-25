package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kynrai/nhshackday-24/ui"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

func (s *Server) handlePageSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.SSE()).Render(r.Context(), w)
}

func (s *Server) handleNotify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	addr := r.FormValue("addr")
	if addr == "" {
		http.Error(w, "addr query parameter is required", http.StatusBadRequest)
		return
	}
	conn, ok := s.clients[addr]
	if !ok {
		http.Error(w, fmt.Sprintf("No client with address: %s", addr), http.StatusBadRequest)
		return
	}
	buf := &bytes.Buffer{}
	ui.Message("Message from server").Render(r.Context(), buf)
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())
}

func (s *Server) handlePageNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.Notifications()).Render(r.Context(), w)
}

func (s *Server) handlePageConnections(w http.ResponseWriter, r *http.Request) {
	var clients []string
	for k := range s.clients {
		clients = append(clients, k)
	}
	w.Header().Set("Content-Type", "text/html")
	ui.Index(ui.Connections(clients)).Render(r.Context(), w)
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	log.Println("Client connected:", r.RemoteAddr)
	s.clients[r.RemoteAddr] = connection

	for {
		mt, message, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break // Exit the loop if the client tries to close the connection or the connection with the interrupted client
		}
		connection.WriteMessage(websocket.TextMessage, message)
	}
	connection.Close()
	delete(s.clients, r.RemoteAddr)
}
