package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres" // ...
	server "iKnowThisWord/internal"
)

// NewDB ...
func NewDB(conf *server.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		conf.InstanceConnectionName,
		conf.DBName,
		conf.Username,
		conf.Password,
	)

	db, err := sql.Open(conf.DriverName, dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
