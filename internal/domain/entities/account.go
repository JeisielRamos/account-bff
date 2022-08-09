package entities

import "time"

type Account struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CPF        int       `json:"cpf"`
	Secret     string    `json:"secret"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}
