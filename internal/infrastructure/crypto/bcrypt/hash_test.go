package bcrypt

import "testing"

func TestGenerateHash(t *testing.T) {
	secret := "senha123"

	hash, err := GenerateHash(secret)

	if err != nil || !CheckSecretHash(secret, hash) {
		t.Errorf("failed to generate hash secret")
	}
}

func TestCheckSecretHash(t *testing.T) {
	secret := "senha123"
	hash := "$2a$14$e4AoZHn4gelDr0eGuivaie50fg90WKJdugMQdvzEr8Lz4MrkW4Im6"
	if !CheckSecretHash(secret, hash) {
		t.Errorf("failed to Check secret hash ")
	}
}
