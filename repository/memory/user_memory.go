package memory

import (
	"control_users/model"
)

type MemoryUserRepository struct {
	users []model.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: []model.User{},
	}
}
func (m *MemoryUserRepository) AddUser(user model.User) {
	m.users = append(m.users, user)
}
func (m *MemoryUserRepository) FilterUsersByStatus() []string {
	result := make([]string, 0)
	for i := range m.users {
		if m.users[i].Status {
			result = append(result, m.users[i].Name)
		}
	}
	return result
}
func (m *MemoryUserRepository) FindUsersByName(name string) (*model.User, error) {
	var err error
	for i := range m.users {
		if m.users[i].Name == name {
			return &m.users[i], nil
		}
	}

	return nil, err
}
func (m *MemoryUserRepository) UpdateUserStatus(name string, status bool) error {
	var err error
	for i := range m.users {
		if m.users[i].Name == name {
			m.users[i].Status = status
			return nil
		}
	}
	return err
}
func (m *MemoryUserRepository) ListUsers() []model.User {
	result := make([]model.User, len(m.users))
	copy(result, m.users)
	return result
}
func (m *MemoryUserRepository) DeleteUser(name string) error {
	var err error
	for i, j := range m.users {
		if j.Name == name {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return err
}
