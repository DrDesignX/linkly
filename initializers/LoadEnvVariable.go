package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return err
	}

	return nil
}
