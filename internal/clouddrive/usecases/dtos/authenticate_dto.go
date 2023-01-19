package dtos

import (
	"time"
)

type DTOAuthenticateReq struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type DTOAuthenticateRes struct {
	AccessToken    string    `json:"access_token"`
	ExpirationDate time.Time `json:"expiration_date"`
}
