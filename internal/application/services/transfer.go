package services

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/models"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/repositories"
)

type TransferServices struct {
	AccountRepository  *repositories.AccountRepository
	TransferRepository *repositories.TransferRepository
}

func NewTransferServices() *TransferServices {
	return &TransferServices{
		repositories.NewAccountRepository(),
		repositories.NewtransferRepository(),
	}
}

func (service *TransferServices) TransferAccountToAccount(cpfUser string, transfer *entities.Transfer) (*entities.Transfer, *entities.Errors) {
	// Get Account Origin
	accountOrigin, err := service.AccountRepository.GetFromCPF(cpfUser)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}
	// Get Account Destination
	accountDestination, err := service.AccountRepository.GetFromAccountID(transfer.Account_destination_id)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	// checks if the balance of the origin account is sufficient for transfer
	balance, _ := strconv.ParseFloat(accountOrigin.Balance, 64)
	if transfer.Amount > balance {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: "Insufficient balance to make the transfer"}
	}

	transfer.Account_origin_id = accountOrigin.ID

	// update account balance
	newBalanceOrigin := balance - transfer.Amount
	accountOrigin.Balance = fmt.Sprintf("%f", newBalanceOrigin)
	_, err = service.AccountRepository.UpdateBalanceFromAccount(accountOrigin)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	balance, _ = strconv.ParseFloat(accountDestination.Balance, 64)
	newBalanceDestination := balance + transfer.Amount
	accountDestination.Balance = fmt.Sprintf("%f", newBalanceDestination)
	_, err = service.AccountRepository.UpdateBalanceFromAccount(accountDestination)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	//Create Transfer
	tsfModels := new(models.TransferModels)
	transferModels := tsfModels.EntitiesToModels(transfer)

	transferResp := new(models.Transfer)
	transferResp, err = service.TransferRepository.Create(transferModels)
	if err != nil {
		return nil, &entities.Errors{StatusCode: fiber.StatusBadRequest, Message: err.Error()}
	}

	return tsfModels.ModelsToEntities(transferResp), nil
}

func (service *TransferServices) Get() {

}
