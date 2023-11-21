package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"api.devbook/src/auth"
	"api.devbook/src/database"
	"api.devbook/src/model"
	"api.devbook/src/repository"
	"api.devbook/src/response"
	"github.com/gorilla/mux"
)

// CreatePublication adiciona uma nova publicação no banco de dados
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication model.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = id

	if err = publication.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	publicationID, err := repo.Create(publication)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	publication.ID = publicationID

	response.JSON(w, http.StatusCreated, publication)
}

// GetAllPublications traz as publicações dos usuários seguidos, dos seguidores e as próprias publicações
func GetAllPublications(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	publications, err := repo.GetAll(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

// GetPublication traz a publicação com base no id fornecido
func GetPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	publication, err := repo.GetById(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

// UpdatePublication atualiza a publicação com base no id fornecido
func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)

	publicationInDB, err := repo.GetById(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationInDB.AuthorID != id {
		response.Error(w, http.StatusForbidden, errors.New("Você não pode editar uma publicação que não pertence à você"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var publication model.Publication
	if err = json.Unmarshal(body, &publication); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.Update(publicationID, publication); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// DeletePublication exclui a publicação com base no id fornecido
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	id, err := auth.ExtractUserID(r)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicationID, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)

	publicationInDB, err := repo.GetById(publicationID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	if publicationInDB.AuthorID != id {
		response.Error(w, http.StatusForbidden, errors.New("Você não pode excluir uma publicação que não pertence à você"))
		return
	}

	if err = repo.Delete(publicationID); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// GetAllPublicationsOfUser retorna todas as publicações de um usuário
func GetAllPublicationsOfUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	authorId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	publications, err := repo.GetAllPublicationsOfUser(authorId)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

// LikePublication incrementa em um a quantidade de curtidas de publicação
func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	if err := repo.Like(publicationId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

// LikePublication decrementa em um a quantidade de curtidas de publicação
func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryOfPublications(db)
	if err := repo.Dislike(publicationId); err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
