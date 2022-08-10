package models

import (
	"fmt"
	"strconv"

	"gitlab.com/desafio-stone/account-bff/internal/domain/entities"
)

type TransferModels struct{}

type Transfer struct {
	ID                     string
	Account_origin_id      string
	Account_destination_id string
	Amount                 string
	Created_at             string
}

func (m *TransferModels) EntitiesToModels(transfer *entities.Transfer) *Transfer {
	amount := fmt.Sprintf("%f", transfer.Amount)
	return &Transfer{
		"",
		transfer.Account_origin_id,
		transfer.Account_destination_id,
		amount,
		"",
	}
}

func (m *TransferModels) ModelsToEntities(transfer *Transfer) *entities.Transfer {
	id, _ := strconv.Atoi(transfer.ID)
	amount, _ := strconv.ParseFloat(transfer.Amount, 64)
	return &entities.Transfer{
		ID:                     id,
		Account_origin_id:      transfer.Account_origin_id,
		Account_destination_id: transfer.Account_destination_id,
		Amount:                 amount,
		Created_at:             transfer.Created_at,
	}
}
