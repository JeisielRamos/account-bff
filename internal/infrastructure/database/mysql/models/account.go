package models

import (
	"fmt"
	"strconv"

	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type AccountModels struct {
	ID         string
	Name       string
	CPF        string
	Secret     string
	Balance    string
	Created_at string
}

func EntitiesToModels(account *entities.Account) *AccountModels {
	balance := fmt.Sprintf("%f", account.Balance)
	return &AccountModels{
		"",
		account.Name,
		account.CPF,
		account.Secret,
		balance,
		"",
	}
}

func ModelsToEntitiesAccount(account *AccountModels) *entities.Account {
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

func ModelsToEntitiesAccountBalance(account *AccountModels) *entities.AccountBalance {
	id, _ := strconv.Atoi(account.ID)
	balance, _ := strconv.ParseFloat(account.Balance, 64)
	return &entities.AccountBalance{
		AccountId: id,
		Balance:   balance,
	}
}
