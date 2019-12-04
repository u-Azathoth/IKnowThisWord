package server_test

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	server "iKnowThisWord/internal"
	"iKnowThisWord/internal/model"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var s *server.Server = &server.Server{}
var store = &postgres.Store{}

func TestMain(m *testing.M) {
	var err error

	absPath, _ := filepath.Abs("../.env.test")
	if err := godotenv.Load(absPath); err != nil {
		log.Fatal(err)
	}

	conf := server.NewConfig()
	Database(conf.DBConfig)

	staticPath := "web/client"
	s, err = server.New(store, &staticPath)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func Database(conf *server.DatabaseConfig) {

	fmt.Println(conf)
	fmt.Println(conf.DBName)

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.InstanceConnectionName,
		conf.DBName,
		conf.Username,
		conf.Password,
	)

	db, err := sql.Open(conf.DriverName, dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	store.DB = db
}

func refreshCardTable() error {
	fmt.Println("1")

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

	fmt.Println("2")

	cards, err := store.Card().Find()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before write")
	fmt.Println(len(cards))

	for i := 0; i < count; i++ {
		c := testCard(i)
		err = store.Card().Save(c)
		if err != nil {
			return err
		}

		fmt.Println("Write")
		fmt.Println(c)
	}

	cards, err = store.Card().Find()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After write")
	fmt.Println(len(cards))

	for _, c := range cards {
		fmt.Println(c)
	}

	return nil
}

func seedCard() (int, error) {
	err := refreshCardTable()
	if err != nil {
		return 0, err
	}

	c := testCard(1)
	err = store.Card().Save(c)
	if err != nil {
		return 0, err
	}

	return c.ID, nil
}

func testCard(uniqueID ...int) *model.Card {
	return &model.Card{
		Word:    fmt.Sprint("consider", uniqueID),
		Meaning: "deem to be",
	}
}
