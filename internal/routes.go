package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ConfigureRouter ...
func (s *Server) ConfigureRouter(staticPath *string) {
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.setupRoutes(staticPath)
}

func (s *Server) setupRoutes(staticPath *string) {
	api := s.Router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/cards", s.HandleCardFind()).Methods("GET")
	api.HandleFunc("/cards/{id}", s.handleCardFindByID()).Methods("GET")
	api.HandleFunc("/cards/{id}", s.handleCardDelete()).Methods("DELETE")
	api.HandleFunc("/cards", s.HandleCardSave()).Methods("POST")

	s.Router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir(*staticPath))),
	)
}
