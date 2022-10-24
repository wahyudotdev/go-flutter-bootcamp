package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			return errors.New("invalid username/password")
		}
	}
	return err
}
