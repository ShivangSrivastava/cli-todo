package config

import (
	"os"

	"github.com/ShivangSrivastava/cli-todo/internal/models"
)

// Load dotenv file
func GetConfig() models.Config {
	config := models.Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
	}

	return config
}
