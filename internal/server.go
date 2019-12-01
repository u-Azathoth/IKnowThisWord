package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"iKnowThisWord/internal/store"
	"net/http"
)

// Server ...
type Server struct {
	Router *mux.Router
	store  store.Store
}

// New ...
func New(store store.Store, staticPath *string) (*Server, error) {
	s := &Server{
		Router: mux.NewRouter(),
		store:  store,
	}

	s.ConfigureRouter(staticPath)

	return s, nil
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
