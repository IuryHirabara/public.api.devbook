package repository

import (
	"database/sql"

	"api.devbook/src/model"
)

// Publications representa um repositório de publicações
type Publications struct {
	db *sql.DB
}

// NewRepositoryOfPublications cria um repositório de publicações
func NewRepositoryOfPublications(db *sql.DB) *Publications {
	return &Publications{db}
}

// Create cria uma publicação no banco de dados
func (repo Publications) Create(publication model.Publication) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO publications (title, content, authorId) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	publicationID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(publicationID), nil
}

// GetById traz a publicação com base no id fornecido
func (repo Publications) GetById(publicationID uint64) (model.Publication, error) {
	row, err := repo.db.Query(
		"SELECT p.*, u.nick FROM publications AS p INNER JOIN users AS u ON P.authorId = u.id WHERE p.id = ?",
		publicationID,
	)

	if err != nil {
		return model.Publication{}, err
	}
	defer row.Close()

	var publication model.Publication
	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return model.Publication{}, err
		}
	}

	return publication, nil
}

// GetAll retorna todas as publicações dos seguidores, dos usuários seguidos e as próprias publicações
func (repo Publications) GetAll(id uint64) ([]model.Publication, error) {
	rows, err := repo.db.Query(
		`SELECT DISTINCT p.*, u.nick FROM publications AS p INNER JOIN users AS u ON p.authorId = u.id
		INNER JOIN followers AS f ON p.authorId = f.userId OR p.authorId = f.followerId
		WHERE f.userId = ? OR f.followerId = ? ORDER BY p.id DESC`,
		id,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []model.Publication
	for rows.Next() {
		var publication model.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Update atualiza uma publicação no banco de dados
func (repo Publications) Update(publicationID uint64, publication model.Publication) error {
	statement, err := repo.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

// Delete exclui uma publicação do banco de dados
func (repo Publications) Delete(publicationID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM publications WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// GetAllPublicationsOfUser retorna todas as publicações de um usuário
func (repo Publications) GetAllPublicationsOfUser(authorId uint64) ([]model.Publication, error) {
	rows, err := repo.db.Query(
		`SELECT p.*, u.nick FROM publications AS p
		INNER JOIN users AS u ON p.authorId = u.id
		WHERE p.authorId = ?`,
		authorId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []model.Publication
	for rows.Next() {
		var publication model.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Like adiciona em um o número de curtidas no banco de dados
func (repo Publications) Like(publicationId uint64) error {
	statement, err := repo.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}

// Dislike subtrai em um o número de curtidas no banco de dados
func (repo Publications) Dislike(publicationId uint64) error {
	statement, err := repo.db.Prepare(
		`UPDATE publications SET likes =
		CASE WHEN likes > 0 THEN likes - 1
		ELSE likes END
		WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationId); err != nil {
		return err
	}

	return nil
}
