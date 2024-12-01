package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TONCONSOLE_KEY string
}

// NewConfig config constructor
func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, err
	}
	config := &Config{
		TONCONSOLE_KEY: os.Getenv("TONCONSOLE_KEY"),
	}
	return config, nil
}
