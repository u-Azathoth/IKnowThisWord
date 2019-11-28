package main

import (
	server "eng-cards/internal"
	"eng-cards/internal/store/postgres"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/env.toml", "path to a config file")
}

func main() {
	flag.Parse()

	// Config
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	CheckError(err)

	// Database
	db, err := postgres.NewDB(config.DatabaseUrl)
	CheckError(err)
	defer db.Close()

	// Server and Store
	store := postgres.NewStore(db)
	s, _ := server.New(store)

	err = http.ListenAndServe(config.BindAddr, s)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
