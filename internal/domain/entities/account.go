package entities

type Account struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	CPF     string  `json:"cpf"`
	Secret  string  `json:"secret"`
	Balance float64 `json:"balance"`
}

type AccountBalance struct {
	AccountId int     `json:"account_id"`
	Balance   float64 `json:"balance"`
}
