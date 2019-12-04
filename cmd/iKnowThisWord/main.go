package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	server "iKnowThisWord/internal"
	"iKnowThisWord/internal/store/postgres"
	"log"
	"net/http"
)

var (
	staticPath string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&staticPath, "static-path", "web/build", "path to a web dist folder")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	db, err := postgres.NewDB(config.DBConfig)
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
