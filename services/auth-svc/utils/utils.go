package utils

import "golang.org/x/crypto/bcrypt"

func CreateHashing(text string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return hashedPassword, err
}
