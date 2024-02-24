package server

func (s *Server) Routes() {
	s.r.HandleFunc("GET /", s.handlePing)
	s.r.HandleFunc("POST /", s.handleSubmit)
	s.r.HandleFunc("GET /ws", s.handleWS)
	s.r.HandleFunc("GET /sms", s.handleSendSMS)
	s.r.HandleFunc("GET /dummy", s.handleDummy)
}
