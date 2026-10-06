package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// AppConfig holds HTTP server configuration
type AppConfig struct {
	Port     string
	Env      string
	LogLevel string
}

// DatabaseConfig holds PostgreSQL connection parameters
type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// Config is the root application configuration
type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

// Function to get environment variables
func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// Load reads configuration from environment variables
func Load(envFile string) (*Config, error) {
	// Load .env file if it exists, ignore error if it doesn't
	_ = godotenv.Load(envFile)

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("config: invalid DB_PORT: %w", err)
	}

	maxOpen, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", ""))
	maxIdle, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", ""))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME_SEC", "300"))

	return &Config{
		App: AppConfig{
			Port:     getEnv("PORT", "8080"),
			Env:      getEnv("ENV", "development"),
			LogLevel: getEnv("LOG_LEVEL", "info"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            dbPort,
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			DBName:          getEnv("DB_NAME", "social_network_db"),
			SSLMode:         getEnv("DB_SSL_MODE", "disable"),
			MaxOpenConns:    maxOpen,
			MaxIdleConns:    maxIdle,
			ConnMaxLifetime: time.Duration(connMaxLifetime) * time.Second,
		},
	}, nil
}
