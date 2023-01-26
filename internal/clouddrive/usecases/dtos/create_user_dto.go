package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type DTOCreateUserReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Pass      string `json:"pass"`
}

type DTOCreateUserRes struct {
	Id        primitive.ObjectID `json:"id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Email     string             `json:"email"`
	Pass      string             `json:"pass"`
}
