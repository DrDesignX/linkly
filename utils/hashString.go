package utils

import (
	"log"

	b "golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string, error) {
	hashPassword, err := b.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		log.Println("Error generating hash password: ", err)
		return "", err

	}
	return string(hashPassword), nil
}

func ComparePassword(hashPassword string, password string) error {
	return b.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
