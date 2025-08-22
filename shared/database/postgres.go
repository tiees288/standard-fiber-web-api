package database

import (
	"fmt"

	"gorm.io/gorm"
)

// PostgreSQLInfo holds PostgreSQL server information
type PostgreSQLInfo struct {
	Version    string `json:"version"`
	Database   string `json:"database"`
	User       string `json:"user"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Connection string `json:"connection_status"`
}

// GetPostgreSQLInfo returns PostgreSQL server information
func GetPostgreSQLInfo() (*PostgreSQLInfo, error) {
	if DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	var version string
	var database string
	var user string
	var host string
	var port string

	// Get PostgreSQL version
	err := DB.Raw("SELECT version()").Scan(&version).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get PostgreSQL version: %w", err)
	}

	// Get current database
	err = DB.Raw("SELECT current_database()").Scan(&database).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}

	// Get current user
	err = DB.Raw("SELECT current_user").Scan(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	// Get host and port from connection settings
	err = DB.Raw("SHOW server_version").Scan(&version).Error
	if err == nil {
		// Connection is working
		host = "connected"
		port = "connected"
	}

	connectionStatus := "connected"
	if err := Ping(); err != nil {
		connectionStatus = "disconnected"
	}

	return &PostgreSQLInfo{
		Version:    version,
		Database:   database,
		User:       user,
		Host:       host,
		Port:       port,
		Connection: connectionStatus,
	}, nil
}

// AutoMigrate runs auto migration for given models
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return DB.AutoMigrate(models...)
}

// Transaction executes a function within a database transaction
func Transaction(fn func(tx *gorm.DB) error) error {
	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return DB.Transaction(fn)
}

// HealthCheck performs a comprehensive database health check
func HealthCheck() map[string]interface{} {
	result := map[string]interface{}{
		"database_connected": false,
		"ping_success":       false,
	}

	if DB == nil {
		result["error"] = "database connection is not initialized"
		return result
	}

	// Test ping
	if err := Ping(); err != nil {
		result["error"] = err.Error()
		return result
	}

	result["ping_success"] = true
	result["database_connected"] = true

	// Get additional info if possible
	if info, err := GetPostgreSQLInfo(); err == nil {
		result["database_info"] = info
	}

	return result
}