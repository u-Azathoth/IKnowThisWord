package server

import (
	"encoding/json"
	"iKnowThisWord/internal/model"
	"net/http"
)

func (s *Server) handleCardFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cards, err := s.store.Card().Find()
		// TODO: Add the custom type of error for NotFoundError
		checkHttpError(err, w)

		s.respond(w, r, http.StatusOK, cards)
	}
}

func (s *Server) handleCardFindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("The get all cards functionality hasn't implemented yet"))

		checkHttpError(err, w)
	}
}

func (s *Server) handleCardSave() http.HandlerFunc {
	type request struct {
		Word    string `json:"word"`
		Meaning string `json:"meaning"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		c := &model.Card{
			Word:    req.Word,
			Meaning: req.Meaning,
		}

		if err := c.Validate(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.Card().Save(c); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, c.ID)
	}
}

// TODO: Remove
func checkHttpError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, "Something went wrong", 500)
	}
}
