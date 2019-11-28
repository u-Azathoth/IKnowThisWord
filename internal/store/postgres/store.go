package postgres

import (
	"database/sql"
	"eng-cards/internal/store"
)

type Store struct {
	db             *sql.DB
	userRepository *CardRepository
}

func (s *Store) Card() store.CardRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &CardRepository{
		store: s,
	}

	return s.userRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
