package server

func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/cards", s.handleCardFind()).Methods("GET")
	s.router.HandleFunc("/cards/{id}", s.handleCardFindById()).Methods("GET")
	s.router.HandleFunc("/cards", s.handleCardSave()).Methods("POST")
}
