package model

import (
	"errors"
	"strings"
	"time"
)

// Publication representa uma publicação feita por um usuário
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare irá chamar os métodos de validação e formatação da publicação
func (publication *Publication) Prepare() error {
	if err := publication.Validate(); err != nil {
		return err
	}

	publication.Format()

	return nil
}

// Validate verifica se os campos estão preenchidos
func (publication *Publication) Validate() error {
	if publication.Title == "" {
		return errors.New("O campo de título deve ser preenchido")
	}

	if publication.Content == "" {
		return errors.New("O campo de conteúdo deve ser preenchido")
	}

	return nil
}

// Format retira os espaços das extremidades dos campos
func (publication *Publication) Format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
