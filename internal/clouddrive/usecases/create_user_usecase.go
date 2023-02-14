package usecases

import (
	"errors"

	"github.com/italoservio/clouddrive/internal/clouddrive/db/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
)

func CreateUser(payload dtos.DTOCreateUserReq) (*dtos.DTOCreateUserRes, error) {
	user, err := repositories.UserByEmail(payload.Email)
	if err != nil {
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	if user != nil {
		return nil, errors.New(custom_errors.BAD_CONFLICT)
	}

	// TODO: Encrypt password...

	user, err = repositories.CreateUser(payload)
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
