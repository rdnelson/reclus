package datamodel

type User struct {
	ID   int
	Name string

	Email    string
	Password string
}

type UserRepo interface {
	GetUser(key string) (*User, error)
	AddUser(key string, user *User) error
	UpdateUser(key string, user *User) error
}
