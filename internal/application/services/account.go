package services

import (
	"github.com/gofiber/fiber/v2"

	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/bcrypt"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/models"

	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/repositories"
)

type AccountServices struct {
	AccountRepository *repositories.AccountRepository
}

func NewAccountServices() *AccountServices {
	return &AccountServices{
		repositories.NewAccountRepository(),
	}
}

func (service *AccountServices) CreateAccount(account *entities.Account) (*entities.Account, *entities.Errors) {
	if len(account.Name) < 1 {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "name cannot be empty"}
	}

	if account.Balance < 0 {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "balance cannot be negative"}
	}

	hash, err := bcrypt.GenerateHash(account.Secret)
	if err != nil || !bcrypt.CheckSecretHash(account.Secret, hash) {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "failed to generate hash secret"}
	}
	account.Secret = hash

	actModels := new(models.AccountModels)
	accountModels := actModels.EntitiesToModels(account)

	accountResp, err := service.AccountRepository.Create(accountModels)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	return actModels.ModelsToEntitiesAccount(accountResp), nil
}

func (service *AccountServices) GetAllAccount() ([]*entities.Account, *entities.Errors) {

	accountsResp, err := service.AccountRepository.GetAll()
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	actModels := new(models.AccountModels)
	accounts := make([]*entities.Account, 0)
	for _, account := range accountsResp {
		accounts = append(accounts, actModels.ModelsToEntitiesAccount(account))
	}

	return accounts, nil
}

func (service *AccountServices) GetBalenceFromAccountID(accountId string) (*entities.AccountBalance, *entities.Errors) {

	accountResp, err := service.AccountRepository.GetFromAccountID(accountId)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	actModels := new(models.AccountModels)
	return actModels.ModelsToEntitiesAccountBalance(accountResp), nil
}
