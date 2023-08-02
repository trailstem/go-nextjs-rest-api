package repository

import (
	"github.com/trailstem/go-nextjs-rest-api/model"
)

// user_repositroy用interface
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
