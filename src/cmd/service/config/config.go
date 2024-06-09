package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Port     string `env:"PORT,default=17001"`
	Secret   string `env:"SECRET,default=SECRET"`
}

type DatabaseConfig struct {
	URL string `env:"DATABASE_URL,default=localhost:5432"`
}

func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Database: DatabaseConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
	}, nil
}
