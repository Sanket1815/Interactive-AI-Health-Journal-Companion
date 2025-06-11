package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	JWTSecret      string
	OpenAIAPIKey   string
	ServerPort     string
	AllowedOrigins []string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "host=localhost port=5432 user=postgres password=1998sanket dbname=mental_health_journal sslmode=disable"),
		JWTSecret:    getEnv("JWT_SECRET", "your_secret_key_change_in_production"),
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		AllowedOrigins: []string{
			getEnv("FRONTEND_URL", "http://localhost:8081"),
		},
	}

	if config.OpenAIAPIKey == "" {
		log.Println("Warning: OPENAI_API_KEY not set")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}