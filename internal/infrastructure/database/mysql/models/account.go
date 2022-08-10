package models

import (
	"fmt"
	"strconv"

	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type AccountModels struct{}

type Account struct {
	ID         string
	Name       string
	CPF        string
	Secret     string
	Balance    string
	Created_at string
}

func (m *AccountModels) EntitiesToModels(account *entities.Account) *Account {
	balance := fmt.Sprintf("%f", account.Balance)
	return &Account{
		"",
		account.Name,
		account.CPF,
		account.Secret,
		balance,
		"",
	}
}

func (m *AccountModels) ModelsToEntitiesAccount(account *Account) *entities.Account {
	id, _ := strconv.Atoi(account.ID)
	balance, _ := strconv.ParseFloat(account.Balance, 64)
	return &entities.Account{
		ID:      id,
		Name:    account.Name,
		CPF:     account.CPF,
		Secret:  account.Secret,
		Balance: balance,
	}
}

func (m *AccountModels) ModelsToEntitiesAccountBalance(account *Account) *entities.AccountBalance {
	id, _ := strconv.Atoi(account.ID)
	balance, _ := strconv.ParseFloat(account.Balance, 64)
	return &entities.AccountBalance{
		AccountId: id,
		Balance:   balance,
	}
}
