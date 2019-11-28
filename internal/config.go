package server

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
