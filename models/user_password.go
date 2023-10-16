package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashed)
}

func ComparePassword(hashed string, normal string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(normal))
	if err != nil {
		return fmt.Errorf("password don't match")
	}
	return nil
}
