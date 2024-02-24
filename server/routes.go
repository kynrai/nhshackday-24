package server

func (s *Server) Routes() {
	s.r.HandleFunc("GET /ping", s.handlePing)
	s.r.HandleFunc("POST /", s.handleSubmit)
	s.r.HandleFunc("GET /ws", s.handleWS)
}
