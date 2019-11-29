package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"iKnowThisWord/internal/store"
	"net/http"
)

type Server struct {
	router *mux.Router
	store  store.Store
}

func New(store store.Store) (*Server, error) {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.SetupRoutes()

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
