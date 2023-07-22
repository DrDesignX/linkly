package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariable() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return err
	}
	return nil
}
