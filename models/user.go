package models

import (
	"errors"

	db "github.com/pageton/authify/db/model"
)

type UserModel struct {
	*db.User
}

type Validate interface {
	Validate() error
}

func (u *UserModel) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
