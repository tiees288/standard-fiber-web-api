package config

import (
	"os"
	"strconv"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
	TimeZone     string
	MaxOpenConns int
	MaxIdleConns int
}

// GetDatabaseConfig returns database configuration from environment variables
func GetDatabaseConfig() *DatabaseConfig {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "5"))

	return &DatabaseConfig{
		Host:         getEnv("DB_HOST", "localhost"),
		Port:         port,
		User:         getEnv("DB_USER", "postgres"),
		Password:     getEnv("DB_PASSWORD", ""),
		DatabaseName: getEnv("DB_NAME", "standard_fiber_db"),
		SSLMode:      getEnv("DB_SSL_MODE", "disable"),
		TimeZone:     getEnv("DB_TIMEZONE", "UTC"),
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
	}
}

// getEnv gets environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}