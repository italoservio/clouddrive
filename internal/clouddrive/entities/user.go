package entities

type User struct {
	Id        string `bson:"_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Pass      string `bson:"pass"`
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
