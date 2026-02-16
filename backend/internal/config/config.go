package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Server    ServerConfig
	Event     EventConfig
	Processor ProcessorConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
}

// EventConfig holds event generator configuration
type EventConfig struct {
	Interval   time.Duration
	Deadline   time.Duration
	BufferSize int
}

// ProcessorConfig holds event processor configuration
type ProcessorConfig struct {
	WorkDuration time.Duration
}

// Load loads configuration from .env file and environment variables with defaults
func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Event: EventConfig{
			Interval:   50 * time.Millisecond,
			Deadline:   100 * time.Millisecond,
			BufferSize: 100,
		},
		Processor: ProcessorConfig{
			WorkDuration: 50 * time.Millisecond,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
