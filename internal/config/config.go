package config

import (
    "os"
)

// Config holds all configuration for the application
type Config struct {
    Port     string
    MongoURI string
    JWTSecret string
}

var AppConfig Config

// LoadEnv loads environment variables into the AppConfig
func LoadEnv() {
    AppConfig = Config{
        Port:     getEnvOrDefault("PORT", "8080"),
        MongoURI: getEnvOrDefault("MONGO_URI", "mongodb://localhost:27017"),
        JWTSecret: getEnvOrDefault("JWT_SECRET", "your_secret_key_here"),
    }
}

// Helper function to get environment variable with a default value
func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}