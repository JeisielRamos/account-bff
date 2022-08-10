package repositories

import (
	"gitlab.com/desafio-stone/account-bff/internal/infrastructure/database/mysql"
)

type TransferRepository struct {
	instance *mysql.DBRepository
}

func NewtransferRepository() *TransferRepository {
	return &TransferRepository{
		instance: mysql.InitDBRepository(),
	}
}

// func (r *TransferRepository) Insert(rqPessoas *request.PessoasPostRq) (*response.PessoaResponse, error) {
// 	status := "0"
// 	if !rqPessoas.Status {
// 		status = "1"
// 	}

// 	sql := "INSERT INTO people(id, name, email, age, state) "
// 	sql += "VALUES (default, '" + rqPessoas.Nome + "', '" + rqPessoas.Email + "', '" + rqPessoas.Idade + "', '" + status + "'); "

// 	res, err := r.instance.DB.Exec(sql)
// 	if err != nil {
// 		return nil, err
// 	}

// 	lastId, err2 := res.LastInsertId()
// 	if err2 != nil {
// 		return nil, err
// 	}

// 	rsp := new(response.PessoaResponse)
// 	rsp.Id = int(lastId)
// 	rsp.Nome = rqPessoas.Nome
// 	rsp.Email = rqPessoas.Email
// 	rsp.Idade, _ = strconv.Atoi(rqPessoas.Idade)
// 	rsp.Status, _ = strconv.Atoi(status)
// 	return rsp, nil
// }
