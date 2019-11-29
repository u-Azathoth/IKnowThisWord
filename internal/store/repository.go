package store

import "iKnowThisWord/internal/model"

type CardRepository interface {
	Save(*model.Card) error
	Find() ([]*model.Card, error)
	FindById(string) (*model.Card, error)
}
