package postgres_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardRepository_Save(t *testing.T) {
	err := refreshCardTable()
	if err != nil {
		t.Fatal(err)
	}

	c := testCard()
	err = store.Card().Save(c)

	assert.NoError(t, err)
	assert.NotNil(t, c.ID)
}

func TestCardRepository_Find(t *testing.T) {
	err := seedCards(3)
	if err != nil {
		t.Fatal(err)
	}

	cards, err := store.Card().Find()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(cards), 3)
}

func TestCardRepository_Delete(t *testing.T) {
	c, err := seedCard()
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, c.ID)

	n, err := store.Card().Delete(c.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n, 1)
}
