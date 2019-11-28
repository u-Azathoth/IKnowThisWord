package store

import "eng-cards/internal/model"

type CardRepository interface {
	Save(*model.Card) error
	Find() ([]*model.Card, error)
	FindById(string) (*model.Card, error)
}
