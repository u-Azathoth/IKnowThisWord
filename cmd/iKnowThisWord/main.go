package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	server "iKnowThisWord/internal"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"net/http"
)

var (
	configPath string
	staticPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/env.toml", "path to a config file")
	flag.StringVar(&staticPath, "static-path", "web/build", "path to a web dist folder")
}

func main() {
	flag.Parse()

	// Config
	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	CheckError(err)

	// Database
	db, err := postgres.NewDB(config.DatabaseURL)
	CheckError(err)
	defer db.Close()

	// Server and Store
	store := postgres.NewStore(db)
	s, _ := server.New(store, &staticPath)

	err = http.ListenAndServe(config.BindAddr, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
	)(s.Router))

	CheckError(err)
}

// CheckError ...
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
