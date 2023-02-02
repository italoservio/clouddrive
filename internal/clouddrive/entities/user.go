package entities

type User struct {
	Id        string 
	FirstName string 
	LastName  string 
	Email     string 
	Pass      string 
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
