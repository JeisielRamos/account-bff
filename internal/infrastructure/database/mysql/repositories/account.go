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

func (repository *AccountRepository) Create(account *models.Account) (*models.Account, error) {
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

func (repository *AccountRepository) UpdateBalanceFromAccount(account *models.Account) (*models.Account, error) {

	sql := `UPDATE accounts	SET balance='` + account.Balance + `' WHERE id= ` + account.ID + `;`

	_, err := repository.instance.DB.Exec(sql)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *AccountRepository) GetAll() ([]*models.Account, error) {
	sql := `SELECT * FROM accounts `
	results, err := repository.instance.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	accounts := make([]*models.Account, 0)
	for results.Next() {
		account := new(models.Account)

		err = results.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.Created_at)
		if err != nil {
			continue
		}
		accounts = append(accounts, account)
	}
	results.Close()

	return accounts, nil
}

func (repository *AccountRepository) GetFromAccountID(accountId string) (*models.Account, error) {
	sql := `SELECT * FROM accounts where id=` + accountId
	results, err := repository.instance.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	account := new(models.Account)

	for results.Next() {
		err = results.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.Created_at)
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repository *AccountRepository) GetFromCPF(cpf string) (*models.Account, error) {
	sql := `SELECT * FROM accounts where cpf='` + cpf + `' `
	results, err := repository.instance.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	account := new(models.Account)

	for results.Next() {
		err = results.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.Created_at)
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}
