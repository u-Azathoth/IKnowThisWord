package store

import "iKnowThisWord/internal/model"

// CardRepository ...
type CardRepository interface {
	Find() ([]*model.Card, error)
	FindByID(string) (*model.Card, error)
	Save(*model.Card) error
	Delete(int) (int, error)
}
