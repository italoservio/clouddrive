package entities

type User struct {
	Id        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Pass      string `json:"pass" bson:"pass"`
}

func NewUser(
	id string,
	first_name string,
	last_name string,
	email string,
	pass string,
) *User {
	return &User{
		Id:        id,
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Pass:      pass,
	}
}
