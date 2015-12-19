package datamodel

type User struct {
	ID       int
	Username string
	Name     string
	Email    string `sql:"index:idx_email_password"`
	Password string `sql:"index:idx_email_password"`
}

type UserRepo interface {
	GetUser(username string) (*User, error)
	GetPartialUser(username string) (*User, error)
	AddUser(user *User) error
	UpdateUser(user *User) error
}
