package dtos

type DTOGetUserByIdReq struct {
	Id string
}

type DTOGetUserByIdRes struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
