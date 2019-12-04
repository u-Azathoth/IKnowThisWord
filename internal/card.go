package server

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"iKnowThisWord/internal/model"
	"net/http"
	"strconv"
)

func (s *Server) HandleCardFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cards, err := s.store.Card().Find()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusOK, cards)
	}
}

func (s *Server) handleCardFindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("The get all cards functionality hasn't implemented yet"))

		if err != nil {
			panic(err)
		}
	}
}

func (s *Server) HandleCardSave() http.HandlerFunc {
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

func (s *Server) HandleCardDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		n, err := s.store.Card().Delete(id)

		if err == nil {
			if n > 0 {
				s.respond(w, r, http.StatusOK, n)
				return
			}

			s.error(w, r, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound)))
			return
		}

		s.error(w, r, http.StatusUnprocessableEntity, err)
	}
}
