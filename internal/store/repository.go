package store

import "iKnowThisWord/internal/model"

type CardRepository interface {
	Find() ([]*model.Card, error)
	FindById(string) (*model.Card, error)
	Save(*model.Card) error
	Delete(int) (int, error)
}
