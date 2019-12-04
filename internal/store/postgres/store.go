package postgres

import (
	"database/sql"
	"iKnowThisWord/internal/store"
)

// Store ...
type Store struct {
	DB             *sql.DB
	cardRepository *CardRepository
}

// Card ...
func (s *Store) Card() store.CardRepository {
	if s.cardRepository != nil {
		return s.cardRepository
	}

	s.cardRepository = &CardRepository{
		store: s,
	}

	return s.cardRepository
}

// NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{DB: db}
}
