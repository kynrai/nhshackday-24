package server

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kynrai/nhshackday-24/model"
	"github.com/kynrai/nhshackday-24/ui"
)

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

func (s *Server) handleSendWS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	addr := r.FormValue("addr")
	if addr == "" {
		http.Error(w, "addr query parameter is required", http.StatusBadRequest)
		return
	}
	conn, ok := s.clients[addr]
	if !ok {
		http.Error(w, "No client with that address", http.StatusBadRequest)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("Hello, World!"))
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

func (s *Server) handleSubmit(w http.ResponseWriter, r *http.Request) {
	var data model.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Data received"))
	// do something with data
}

func (s *Server) handleSendSMS(w http.ResponseWriter, r *http.Request) {
	// send sms
	resp, err := s.tw.SendSMS("+447927303651", "Twilio", "Hello, World!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(resp)
}

func (s *Server) handleDummy(w http.ResponseWriter, r *http.Request) {
	data, err := s.ian.Read()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func (s *Server) handleClinicianView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ui.Index().Render(r.Context(), w)
}
