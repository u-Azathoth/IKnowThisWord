package postgres_test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"iKnowThisWord/internal/model"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"os"
	"testing"
)

var store = &postgres.Store{}

func TestMain(m *testing.M) {
	Database()
	os.Exit(m.Run())
}

func Database() {
	databaseURL := "host=localhost dbname=eng_cards_test sslmode=disable"
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	store.DB = db
}

func refreshCardTable() error {
	_, err := store.DB.Exec("TRUNCATE card CASCADE")
	if err != nil {
		return err
	}

	return nil
}

func seedCard() (*model.Card, error) {
	err := refreshCardTable()
	if err != nil {
		return nil, err
	}

	c := testCard(0)

	err = store.Card().Save(c)
	if err != nil {
		log.Fatal(err)
	}

	return c, nil
}

func seedCards(count int) error {
	err := refreshCardTable()
	if err != nil {
		return err
	}

	for i := 0; i < count; i++ {
		err = store.Card().Save(testCard(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func testCard(uniqueId ...int) *model.Card {
	return &model.Card{
		Word:    fmt.Sprint("consider", uniqueId),
		Meaning: "deem to be",
	}
}
