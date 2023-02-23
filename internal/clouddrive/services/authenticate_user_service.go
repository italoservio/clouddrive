package services

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/italoservio/clouddrive/internal/clouddrive/db/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(payload dtos.DTOAuthenticateReq) (*dtos.DTOAuthenticateRes, error) {
	user_repository := repositories.NewUserRepository()
	user, err := user_repository.UserByEmail(payload.Email)
	if err != nil {
		return nil, errors.New(custom_errors.BAD_DB_CALL)
	}

	if (user == &entities.User{}) {
		return nil, errors.New(custom_errors.NOT_FOUND_REGISTRY)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(payload.Pass))
	if err != nil {
		return nil, errors.New(custom_errors.NOT_FOUND_REGISTRY)
	}

	expiration_date := time.Now().Add(2 * time.Hour)
	unsigned_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       1,
		"iss":       "clouddrive",
		"iat":       time.Now(),
		"exp":       expiration_date,
		"firstName": user.FirstName,
	})

	access_token, err := unsigned_token.SignedString([]byte("SUPERSECRETO"))
	if err != nil {
		return nil, errors.New(custom_errors.UNEXPECTED_ERR)
	}

	return &dtos.DTOAuthenticateRes{
		AccessToken:    access_token,
		ExpirationDate: expiration_date.UTC(),
	}, nil
}
