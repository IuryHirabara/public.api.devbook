package model

// AuthData contém o id e o token do usuário autenticado
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
