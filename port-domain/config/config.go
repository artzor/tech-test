// Package config provides application configuration
package config

import (
	"os"
)

// Config contains application parameters
type Config struct {
	ServerPort string
	DBInstance string
	DBName     string
}

// Load returns Config instance with fields loaded from environment variables or predefined values
func Load() Config {
	return Config{
		ServerPort: envDefault("SERVICE_PORT", ":9090"),
		DBName:     envDefault("SERVICE_DB_NAME", "portdomain"),
		DBInstance: envDefault("SERVICE_DB_URI", ""),
	}
}

func envDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
