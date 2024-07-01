package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	EnvConfig struct {
		Port      string
		SecretKey SecretKey
	}

	SecretKey string
)

func GetEnvConfig() *EnvConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT should not be empty")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY should not be empty")
	}
	return &EnvConfig{
		Port:      port,
		SecretKey: SecretKey(secretKey),
	}
}
