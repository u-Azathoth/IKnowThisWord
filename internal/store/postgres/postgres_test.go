package postgres_test

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // ...
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	server "iKnowThisWord/internal"
	"path/filepath"

	"iKnowThisWord/internal/model"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"os"
	"testing"
)

var store = &postgres.Store{}

func TestMain(m *testing.M) {
	absPath, _ := filepath.Abs("../../../.env.test")
	if err := godotenv.Load(absPath); err != nil {
		log.Fatal(err)
	}

	conf := server.NewConfig()
	Database(conf.DBConfig)

	os.Exit(m.Run())
}

func Database(conf *server.DatabaseConfig) {
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

func testCard(uniqueID ...int) *model.Card {
	return &model.Card{
		Word:    fmt.Sprint("consider", uniqueID),
		Meaning: "deem to be",
	}
}
