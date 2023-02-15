package services

import (
	"errors"

	"github.com/italoservio/clouddrive/internal/clouddrive/db/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func UserById(payload dtos.DTOGetUserByIdReq) (*dtos.DTOGetUserByIdRes, error) {
	var user *entities.User
	user_repository := repositories.NewUserRepository()
	user, err := user_repository.UserById(payload.Id)
	if err != nil {
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	if (*user == entities.User{}) {
		return nil, errors.New(custom_errors.NOT_FOUND_REGISTRY)
	}

	return &dtos.DTOGetUserByIdRes{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
