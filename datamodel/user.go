package datamodel

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `sql:"index:idx_email_password" json:"email"`
	Password string `sql:"index:idx_email_password" json:"-"`
}

type UserRepo interface {
	GetUser(username string) (*User, error)
	GetPartialUser(username string) (*User, error)
	AddUser(user *User) error
	UpdateUser(user *User) error
}
