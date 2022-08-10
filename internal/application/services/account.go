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

func (service *AccountServices) Create(account *entities.Account) (*entities.Account, *entities.Errors) {
	if len(account.Name) < 1 {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "name cannot be empty"}
	}

	if account.Balance < 0 {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "balance cannot be negative"}
	}

	hash, err := bcrypt.GenerateHash(account.Secret)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}
	account.Secret = hash

	accountModels := models.EntitiesToModels(account)

	accountResp, err := service.AccountRepository.Create(accountModels)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	return models.ModelsToEntities(accountResp), nil
}
