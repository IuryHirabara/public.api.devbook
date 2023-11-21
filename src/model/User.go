package model

import (
	"errors"
	"strings"
	"time"

	"api.devbook/src/security"
	"github.com/badoux/checkmail"
)

// Representa um usuários utilizando a rede social
type User struct {
	// "omitempty" serve para ocultar parâmetros não recebidos para json
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare chama os métodos para validar e formatar os campos
func (u *User) Prepare(stage string) error {
	if err := u.verifyFields(stage); err != nil {
		return err
	}

	if err := u.formatFields(stage); err != nil {
		return err
	}

	return nil
}

func (u *User) verifyFields(stage string) error {
	if u.Name == "" {
		return errors.New("O campo nome deve ser preenchido")
	}

	if u.Nick == "" {
		return errors.New("O campo de apelido deve ser preenchido")
	}

	if u.Email == "" {
		return errors.New("O campo email deve ser preenchido")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if stage == "register" && u.Password == "" {
		return errors.New("O campo senha deve ser preenchido")
	}

	return nil
}

func (u *User) formatFields(stage string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if stage == "register" {
		passwordHashed, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(passwordHashed)
	}

	return nil
}
