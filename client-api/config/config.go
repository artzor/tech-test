// Package config package provides application configuration
package config

import (
	"os"
)

// Config contains application parameters
type Config struct {
	FileName   string
	ServerPort string
	PortDomain string
}

// Load returns Config instance with fields loaded from environment variables or predefined default values
func Load() Config {
	return Config{
		FileName:   envDefault("IMPORT_FILE", ""),
		ServerPort: envDefault("SERVER_PORT", ":8080"),
		PortDomain: envDefault("PORT_DOMAIN", "localhost:9090"),
	}
}

func envDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
