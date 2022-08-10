package repositories

import (
	"strconv"
	"time"

	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/models"
)

type TransferRepository struct {
	instance *mysql.DBRepository
}

func NewtransferRepository() *TransferRepository {
	return &TransferRepository{
		instance: mysql.InitDBRepository(),
	}
}

func (repository *TransferRepository) Create(transfer *models.Transfer) (*models.Transfer, error) {
	transfer.Created_at = time.Now().Format("2006-01-02T15:04:05")
	sql := `INSERT INTO transfers(id, account_origin_id, account_destination_id, amount, created_at) 
			VALUES (default, '` + transfer.Account_origin_id + `', '` + transfer.Account_destination_id + `', '` + transfer.Amount + `', '` + transfer.Created_at + `');`

	res, err := repository.instance.DB.Exec(sql)
	if err != nil {
		return nil, err
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		return nil, err
	}

	transfer.ID = strconv.Itoa(int(lastId))
	return transfer, nil
}
