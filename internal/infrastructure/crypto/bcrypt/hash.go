package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), 14)
	return string(bytes), err
}

func CheckSecretHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
