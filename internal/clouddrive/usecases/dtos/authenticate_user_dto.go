package dtos

type DTOAuthenticateUserReq struct {
	Email     string `json:"email"`
	Pass      string `json:"pass"`
}

type DTOAuthenticateUserRes struct {
	Id        string			 `json:"id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Pass      string             `json:"pass"`
}
