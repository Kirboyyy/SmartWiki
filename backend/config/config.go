package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBPath    string
	OpenAIKey string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables from the system")
	}

	DBPath = os.Getenv("DB_PATH")
	OpenAIKey = os.Getenv("OPENAI_KEY")

	if DBPath == "" || OpenAIKey == "" {
		log.Fatal("Required environment variables DB_PATH or OPENAI_KEY are not set")
	}
}
