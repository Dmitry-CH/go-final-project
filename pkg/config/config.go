package config

import "os"

type Config struct {
	DBfile   string
	Password string
	Port     string
}

func New() *Config {
	conf := Config{
		DBfile:   os.Getenv("TODO_DBFILE"),
		Password: os.Getenv("TODO_PASSWORD"),
		Port:     os.Getenv("TODO_PORT"),
	}

	return &conf
}
