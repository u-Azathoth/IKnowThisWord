package postgres

import (
	"database/sql"
	"iKnowThisWord/internal/model"
	"iKnowThisWord/internal/store"
)

// CardRepository ...
type CardRepository struct {
	store *Store
}

// Find ...
func (c *CardRepository) Find() ([]*model.Card, error) {
	cards := []*model.Card{}

	rows, err := c.store.DB.Query(
		"SELECT id, word, meaning, recognition_rate from card order by id",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		card := &model.Card{}
		err = rows.Scan(&card.ID, &card.Word, &card.Meaning, &card.RecognitionRate)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

// FindByID ...
func (c *CardRepository) FindByID(id string) (*model.Card, error) {
	card := &model.Card{}

	err := c.store.DB.QueryRow(
		"SELECT id, word, meaning, recognition_rate from card where card.id = $1", id,
	).Scan(&card.ID, &card.Word, &card.Meaning, &card.RecognitionRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return card, nil
}

// Save ...
func (c *CardRepository) Save(card *model.Card) error {
	return c.store.DB.QueryRow(
		"INSERT INTO card (word, meaning) VALUES ($1, $2) returning id",
		card.Word,
		card.Meaning,
	).Scan(&card.ID)
}

// Delete ...
func (c *CardRepository) Delete(id int) (int, error) {
	res, err := c.store.DB.Exec("DELETE FROM card WHERE card.id = $1", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
