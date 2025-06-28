package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/lpernett/godotenv"
)

var (
	config *Config
	once   sync.Once
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Env  string
}

// LoadConfig loads configuration from environment variables (singleton pattern)
func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file only once
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using system environment variables")
		}

		config = &Config{
			Database: DatabaseConfig{
				Host:     getEnv("DB_HOST", ""),
				Port:     getEnv("DB_PORT", "3306"),
				User:     getEnv("DB_USER", ""),
				Password: getEnv("DB_PASSWORD", ""), // Empty default
				Name:     getEnv("DB_NAME", "BOOK_STORE"),
			},
			Server: ServerConfig{
				Port: getEnv("APP_PORT", "8080"),
				Env:  getEnv("APP_ENV", "development"),
			},
		}

		// Validate required configuration
		if config.Database.Password == "" {
			log.Fatal("DB_PASSWORD environment variable is required")
		}
	})

	return config
}

// GetDSN returns the database connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}

// IsProduction returns true if the application is running in production
func (c *Config) IsProduction() bool {
	return c.Server.Env == "production"
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt gets an environment variable as int or returns a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
