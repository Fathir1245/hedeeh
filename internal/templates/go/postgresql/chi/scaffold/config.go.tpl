package config

import "os"

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
	AppPort string
}

func Load() Config {
	return Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		AppPort: os.Getenv("APP_PORT"),
	}
}
