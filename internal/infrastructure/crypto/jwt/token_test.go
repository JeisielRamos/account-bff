package jwt

import "testing"

func TestGenerateToken(t *testing.T) {
	cpf := "01816154016"

	_, err := GenerateToken(cpf)
	if err != nil {
		t.Errorf("failed to generate token")
	}
}

func TestValidadeToken(t *testing.T) {
	cpf := "01816154016"

	token, err := GenerateToken(cpf)
	if err != nil {
		t.Errorf("failed to generate token")
	}

	_, err = ValidadeToken(token)
	if err != nil {
		t.Errorf("failed to validade token")
	}
}
