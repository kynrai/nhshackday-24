package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	r  *chi.Mux
	db *gorm.DB
	tw *Twilio
}
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "nshhackday2024" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewServer(conf Config) (*Server, error) {
	db, err := NewDB("data.db")
	if err != nil {
		return nil, err
	}
	r := chi.NewRouter()
	r.Use(auth)
	db.AutoMigrate(&Product{})
	return &Server{
		r:  r,
		tw: NewTwilio(conf.TwilioSID, conf.TwilioAuth, conf.TwilioFrom),
		db: db,
	}, nil
}

func (s *Server) Start() {
	s.Routes()
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", s.r)
}
