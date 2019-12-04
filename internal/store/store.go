package store

// Store ...
type Store interface {
	Card() CardRepository
}
