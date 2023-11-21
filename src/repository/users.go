package repository

import (
	"database/sql"
	"fmt"

	"api.devbook/src/model"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewRepositoryOfUsers Cria um repositório de usuários
func NewRepositoryOfUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Criar insere um usuário no banco de dados
func (repo Users) Create(user model.User) (uint64, error) {
	statement, err := repo.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(userID), nil
}

// Get traz todos os usuários que atendem o filtro
func (repo Users) GetAll(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repo.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID traz o usuário conforme o id fornecido
func (repo Users) GetByID(id uint64) (model.User, error) {
	row, err := repo.db.Query("SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", id)
	if err != nil {
		return model.User{}, err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// Update atualiza as informações de um usuário
func (repo Users) Update(id uint64, user model.User) error {
	statement, err := repo.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

// Exclui o registro de um usuário do banco de dados
func (repo Users) Delete(id uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// SearchByEmail busca um usuário pelo email informado e retorna o seu id e o hash da senha
func (repo Users) SearchByEmail(email string) (model.User, error) {
	row, err := repo.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return model.User{}, err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// Follow permite que um usuário siga outro
func (repo Users) Follow(id, followedID uint64) error {
	statement, err := repo.db.Prepare("INSERT ignore INTO followers (userId, followerId) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(followedID, id); err != nil {
		return err
	}

	return nil
}

// Unfollow permite que um usuário deixe de seguir outro
func (repo Users) Unfollow(id, followedID uint64) error {
	statement, err := repo.db.Prepare("DELETE FROM followers WHERE userId = ? AND followerId = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(followedID, id); err != nil {
		return err
	}

	return nil
}

// Busca os seguidores de um usuário
func (repo Users) GetAllFollowers(id uint64) ([]model.User, error) {
	rows, err := repo.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM followers AS f
		INNER JOIN users AS u ON f.followerId = u.id WHERE userId = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []model.User
	for rows.Next() {
		var follower model.User

		err = rows.Scan(&follower.ID, &follower.Name, &follower.Nick, &follower.Email, &follower.CreatedAt)
		if err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

// GetAllFollowing retorna todos os usuários que o usuário está seguindo conforme o id passado
func (repo Users) GetAllFollowing(id uint64) ([]model.User, error) {
	rows, err := repo.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM followers AS f
		INNER JOIN users AS u ON f.userId = u.id WHERE followerId = ?`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allUsersFollowing []model.User
	for rows.Next() {
		var userFollowing model.User

		err = rows.Scan(
			&userFollowing.ID,
			&userFollowing.Name,
			&userFollowing.Nick,
			&userFollowing.Email,
			&userFollowing.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		allUsersFollowing = append(allUsersFollowing, userFollowing)
	}

	return allUsersFollowing, nil
}

// SearchPasswordByUserID traz a senha de um usuário pelo id fornecido
func (repo Users) SearchPasswordByUserID(id uint64) (string, error) {
	row, err := repo.db.Query("SELECT password FROM users WHERE id = ?", id)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// ChangePassword altera a senha do usuário no banco de dados
func (repo Users) ChangePassword(id uint64, password string) error {
	statement, err := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(password, id); err != nil {
		return err
	}

	return nil
}
