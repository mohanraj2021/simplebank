package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedPasswordByte, berr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if berr != nil {
		return "", fmt.Errorf("failed to hash password %w", berr)
	}
	err := bcrypt.CompareHashAndPassword(hashedPasswordByte, []byte(password))
	fmt.Println(err)
	return string(hashedPasswordByte), nil
}

func CheckPassword(hasedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}
