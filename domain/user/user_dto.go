package user

import (
	"strings"
	"tripplanner/utils/error"
)

const (
	StatusActive = `active`
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"createdAt"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

func (user *User) Validate() *error.AppErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return error.NewBadRequestError("Email not found")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return error.NewBadRequestError("Password not valid")
	}
	return nil
}
