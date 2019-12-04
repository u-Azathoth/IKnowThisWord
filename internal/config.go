package server

import "os"

// DatabaseConfig ...
type DatabaseConfig struct {
	DriverName             string
	InstanceConnectionName string
	DBName                 string
	Username               string
	Password               string
}

// Config ...
type Config struct {
	BindAddr string
	DBConfig *DatabaseConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: getEnv("BIND_ADDR", ":8080"),
		DBConfig: &DatabaseConfig{
			DriverName:             getEnv("DRIVER_NAME", ""),
			InstanceConnectionName: getEnv("INSTANCE_CONNECTION_NAME", ""),
			DBName:                 getEnv("DB_NAME", ""),
			Username:               getEnv("DB_USER", ""),
			Password:               getEnv("DB_PASSWORD", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
