package model

// Password representa o formato da requição de atualização de senha
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
