package usecases

import (
	"errors"

	"github.com/italoservio/clouddrive/internal/clouddrive/db/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func CreateUser(payload dtos.DTOCreateUserReq) (*dtos.DTOCreateUserRes, error) {
	user_repository := repositories.NewUserRepository()
	user, err := user_repository.UserByEmail(payload.Email)
	if err != nil {
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	if (*user != entities.User{}) {
		return nil, errors.New(custom_errors.BAD_CONFLICT)
	}

	// TODO: Encrypt password...

	user, err = user_repository.CreateUser(payload)
	if err != nil {
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	return &dtos.DTOCreateUserRes{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
