package dtos

type DTOGetUserReq struct {
	Email     string `json:"email"`
}

type DTOGetUserRes struct {
	Id        string
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Pass      string             `json:"pass"`
}
