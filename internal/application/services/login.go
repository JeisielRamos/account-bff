package services

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/bcrypt"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/crypto/jwt"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/repositories"
)

type LoginServices struct {
	AccountRepository *repositories.AccountRepository
}

func NewLoginServices() *LoginServices {
	return &LoginServices{
		repositories.NewAccountRepository(),
	}
}

func (service *LoginServices) AuthenticateUser(login *entities.Login) (*entities.UserToken, *entities.Errors) {

	accountsResp, err := service.AccountRepository.GetFromCPF(login.Cpf)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "invalid CPF or Secret"}
	}

	if !bcrypt.CheckSecretHash(login.Secret, accountsResp.Secret) {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "invalid CPF or Secret"}
	}

	token, err := jwt.GenerateToken(login.Cpf)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "failed to generate Token"}
	}

	return &entities.UserToken{Token: token}, nil
}
