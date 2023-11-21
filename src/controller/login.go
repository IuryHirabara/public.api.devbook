package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"api.devbook/src/auth"
	"api.devbook/src/database"
	"api.devbook/src/model"
	"api.devbook/src/repository"
	"api.devbook/src/response"
	"api.devbook/src/security"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfUsers(db)
	userOfDB, err := repo.SearchByEmail(user.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(user.Password, userOfDB.Password); err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(userOfDB.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	authData := model.AuthData{
		ID:    strconv.FormatUint(userOfDB.ID, 10),
		Token: token,
	}

	response.JSON(w, http.StatusOK, authData)
}
