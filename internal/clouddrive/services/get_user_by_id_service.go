package services

import (
	"errors"

	"github.com/italoservio/clouddrive/internal/clouddrive/db/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserById(payload dtos.DTOGetUserByIdReq) (*dtos.DTOGetUserByIdRes, error) {
	var user *entities.User
	user, err := repositories.UserById(payload.Id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(custom_errors.NOT_FOUND_REGISTRY)
		}
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	return &dtos.DTOGetUserByIdRes{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
