package service

import (
	"control_users/model"
	"control_users/repository"
)

type UserService struct {
	repo repository.UserRepositore
}

func NewUserService(repo repository.UserRepositore) *UserService {
	return &UserService{
		repo: repo,
	}
}

// func input name and status for user #1
func (s *UserService) AddUser(u model.User) {
	s.repo.AddUser(u)
}

// filter useres by status
func (s *UserService) FilterUsersByStatus() []string {
	return s.repo.FilterUsersByStatus()
}

// find users by name #3 and #4
func (s *UserService) FindUsersByName(name string) (*model.User, error) {
	res, err := s.repo.FindUsersByName(name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// update user status by name #4
func (s *UserService) UpdateUserStatus(name string, status bool) error {
	if err := s.repo.UpdateUserStatus(name, status); err != nil {
		return err
	}
	return nil
}

// #5 display all users
func (s *UserService) ListUsers() []model.User {
	result := s.repo.ListUsers()
	return result
}

// #6 delete user by name
func (s *UserService) DeleteUser(name string) error {
	if err := s.repo.DeleteUser(name); err != nil {
		return err
	}
	return nil
}
