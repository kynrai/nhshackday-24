package server

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Server struct {
	r  *http.ServeMux
	db *gorm.DB
	tw *Twilio
}
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func NewServer(conf Config) (*Server, error) {
	db, err := NewDB("data.db")
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Product{})
	return &Server{
		r:  http.NewServeMux(),
		tw: NewTwilio(conf.TwilioSID, conf.TwilioAuth, conf.TwilioFrom),
		db: db,
	}, nil
}

func (s *Server) Start() {
	s.Routes()
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", s.r)
}
