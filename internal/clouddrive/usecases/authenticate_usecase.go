package usecases

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	custom_errors "github.com/italoservio/clouddrive/internal/clouddrive/errors"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
)

func Authenticate(payload dtos.DTOAuthenticateReq) (*dtos.DTOAuthenticateRes, error) {
	expiration_date := time.Now().Add(2 * time.Hour)
	unsigned_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": 1,
		"iss": "clouddrive",
		"iat": time.Now(),
		"exp": expiration_date,
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
