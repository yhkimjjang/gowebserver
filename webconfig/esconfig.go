package webconfig

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EsConfig struct {
	EsAddress     string `env:"Address,file"`
	EsUserName    string `env:"Username,file"`
	EsPassword    string `env:"Password,file"`
	EsFingerprint string `env:"Fingerprint,file"`
}

func SetupElasticsearchConfig() EsConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := EsConfig{
		EsAddress:     os.Getenv("Address"),
		EsUserName:    os.Getenv("Username"),
		EsPassword:    os.Getenv("Password"),
		EsFingerprint: os.Getenv("Fingerprint"),
	}

	return cfg
}
