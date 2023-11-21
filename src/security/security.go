package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e retorna um hash dessa string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compara a senha com o hash informado
func VerifyPassword(password, passwordHashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
