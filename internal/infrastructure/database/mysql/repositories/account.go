package repositories

import (
	"strconv"
	"time"

	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql"
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql/models"
)

type AccountRepository struct {
	instance *mysql.DBRepository
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		instance: mysql.InitDBRepository(),
	}
}

func (repository *AccountRepository) Create(account *models.AccountModels) (*models.AccountModels, error) {
	account.Created_at = time.Now().Format("2006-01-02T15:04:05")
	sql := `INSERT INTO accounts(id, name, cpf, secret, balance, created_at) 
			VALUES (default, '` + account.Name + `', '` + account.CPF + `', '` + account.Secret + `', '` + account.Balance + `', '` + account.Created_at + `');`

	res, err := repository.instance.DB.Exec(sql)
	if err != nil {
		return nil, err
	}

	lastId, err2 := res.LastInsertId()
	if err2 != nil {
		return nil, err
	}
	account.ID = strconv.Itoa(int(lastId))
	return account, nil
}
