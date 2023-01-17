package entities

type User struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func CreateUser(full_name string, email string) *User {
	return &User{
		FullName: full_name,
		Email:    email,
	}
}
