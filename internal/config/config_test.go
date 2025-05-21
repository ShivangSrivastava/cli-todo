package config

import (
	"os"
	"testing"

	"github.com/ShivangSrivastava/cli-todo/internal/models"
)

// Test the GetConfig function
func TestGetConfig(t *testing.T) {
	// Set environment variables for the test
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB", "testdb")

	config := GetConfig()

	expected := models.Config{
		PostgresUser:     "testuser",
		PostgresPassword: "testpassword",
		PostgresDB:       "testdb",
	}

	if config != expected {
		t.Errorf("Expected %v, got %v", expected, config)
	}
}
