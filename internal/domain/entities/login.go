package entities

type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type UserToken struct {
	Token string `json:"token"`
}
