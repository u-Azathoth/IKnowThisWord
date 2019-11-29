package server

import "net/http"

func (s *Server) SetupRoutes(staticPath *string) {
	api := s.router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/cards", s.handleCardFind()).Methods("GET")
	api.HandleFunc("/cards/{id}", s.handleCardFindById()).Methods("GET")
	api.HandleFunc("/cards/{id}", s.handleCardDelete()).Methods("DELETE")
	api.HandleFunc("/cards", s.handleCardSave()).Methods("POST")

	s.router.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir(*staticPath))),
	)
}
