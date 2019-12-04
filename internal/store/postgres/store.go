package postgres

import (
	"database/sql"
	"iKnowThisWord/internal/store"
)

type Store struct {
	DB             *sql.DB
	cardRepository *CardRepository
}

func (s *Store) Card() store.CardRepository {
	if s.cardRepository != nil {
		return s.cardRepository
	}

	s.cardRepository = &CardRepository{
		store: s,
	}

	return s.cardRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{DB: db}
}
