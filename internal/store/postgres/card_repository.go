package postgres

import (
	"iKnowThisWord/internal/model"
)

type CardRepository struct {
	store *Store
}

func (c *CardRepository) Find() ([]*model.Card, error) {
	cards := []*model.Card{}

	rows, err := c.store.db.Query("SELECT id, word, meaning from card order by id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		card := &model.Card{}
		err = rows.Scan(&card.ID, &card.Word, &card.Meaning)
		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *CardRepository) FindById(string) (*model.Card, error) {
	panic("implement me")
}

func (c *CardRepository) Save(card *model.Card) error {
	return c.store.db.QueryRow(
		"INSERT INTO card (word, meaning) VALUES ($1, $2) returning id",
		card.Word,
		card.Meaning,
	).Scan(&card.ID)
}

func (c *CardRepository) Delete(id int) (int, error) {
	res, err := c.store.db.Exec("DELETE FROM card WHERE card.id = $1", id)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
