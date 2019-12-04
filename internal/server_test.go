package server_test

import (
	"database/sql"
	"fmt"
	server "iKnowThisWord/internal"
	"iKnowThisWord/internal/model"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"os"
	"testing"
)

var s *server.Server = &server.Server{}
var store = &postgres.Store{}

func TestMain(m *testing.M) {
	var err error

	Database()

	staticPath := "web/client"
	s, err = server.New(store, &staticPath)
	if err != nil {
		log.Fatal(err)
	}

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
