package repository

import "control_users/model"

type UserRepositore interface {
	AddUser(user model.User)
	FilterUsersByStatus() []string
	FindUsersByName(name string) (*model.User, error)
	UpdateUserStatus(name string, status bool) error
	ListUsers() []model.User
	DeleteUser(name string) error
}
