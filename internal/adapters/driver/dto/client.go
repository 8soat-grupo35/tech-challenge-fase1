package dto

type ClientDto struct {
	Id   uint32 `param:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}
