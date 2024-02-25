package server

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kynrai/nhshackday-24/model"
	"github.com/kynrai/nhshackday-24/ui"
	"github.com/kynrai/nhshackday-24/ui/clinician"
	"github.com/kynrai/nhshackday-24/ui/patient"
)

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
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
	resp, err := s.tw.SendSMS(s.tw.to, "Twilio", "Hello, World!")
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
	data, err := s.ian.Read()
	tabType := "prescription"

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	labs := s.reports
	w.Header().Set("Content-Type", "text/html")
	ui.Index(clinician.ClinicianView(tabType, *data, labs)).Render(r.Context(), w)
}

func (s *Server) handleClinicianNav(w http.ResponseWriter, r *http.Request) {
	tabType := chi.URLParam(r, "type")
	if tabType == "" {
		tabType = "prescription"
	}
	data, err := s.ian.Read()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	labs := s.reports
	w.Header().Set("Content-Type", "text/html")
	clinician.ClinicianView(tabType, *data, labs).Render(r.Context(), w)
}

func (s *Server) handlePatientView(w http.ResponseWriter, r *http.Request) {
	tabType := chi.URLParam(r, "type")
	if tabType == "" {
		tabType = "profile"
	}
	data, err := s.ian.Read()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	labs := s.reports

	if r.URL.Query().Get("ack") == "true" {
		fmt.Println("ack")
		buf := new(bytes.Buffer)
		clinician.Ack().Render(r.Context(), buf)
		s.notifyAckMsg <- buf.String()
		w.Header().Set("HX-Push-Url", "/patient")
	}
	w.Header().Set("Content-Type", "text/html")
	patient.PatientIndex(patient.PatientView(tabType, *data, labs)).Render(r.Context(), w)
}

func (s *Server) handlePatientNav(w http.ResponseWriter, r *http.Request) {
	tabType := chi.URLParam(r, "type")
	if tabType == "" {
		tabType = "profile"
	}
	data, err := s.ian.Read()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	labs := s.reports

	w.Header().Set("Content-Type", "text/html")
	patient.PatientView(tabType, *data, labs).Render(r.Context(), w)
}

func (s *Server) handleSMSReminder(w http.ResponseWriter, r *http.Request) {
	// send sms
	resp, err := s.tw.SendSMS(s.tw.to, "Twilio", "Reminder: Your next blood test is due on 23/3/2024. Please call 01234567890 to book an appointment with your GP.")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(resp)
}

func (s *Server) handleSMSAlert(w http.ResponseWriter, r *http.Request) {
	// send sms
	resp, err := s.tw.SendSMS(s.tw.to, "Twilio", `*NHS ALERT*: Your blood test results are not within the normal range. Stop taking your medication IMMEDIATELY and call 01234567890 to book an appointment as soon as possible.`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(resp)
}
