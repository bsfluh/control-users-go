package postgres

import (
	"control_users/model"
	"database/sql"
	"fmt"
	"strings"
)

type DBUserRepositore struct {
	db *sql.DB
}

func NewDBUserRepositore(db *sql.DB) *DBUserRepositore {
	return &DBUserRepositore{db: db}
}

func (m *DBUserRepositore) AddUser(user model.User) {
	query := `
INSERT INTO users (name,status)
VALUES ($1, $2);
`
	_, err := m.db.Exec(query, user.Name, user.Status)
	if err != nil {
		fmt.Println("insert error user: ", err)
		return
	}
}
func (m *DBUserRepositore) FilterUsersByStatus() []string {
	query := `
SELECT name
FROM users
WHERE status=$1;
`
	rows, err := m.db.Query(query, true)
	if err != nil {
		fmt.Println("error found by status")
		return nil
	}
	defer rows.Close()
	users := make([]string, 0)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("error scan name for found by status")
			return nil
		}
		users = append(users, name)
	}
	return users
}
func (m *DBUserRepositore) FindUsersByName(name string) (*model.User, error) {
	if m == nil || m.db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	name = strings.TrimSpace(name)
	query := `
	SELECT name,status
	FROM users
	WHERE name=$1;
	`
	rows := m.db.QueryRow(query, name)
	var user model.User
	err := rows.Scan(&user.Name, &user.Status)
	if err != nil {
		fmt.Println("error found by name")
		if err == sql.ErrNoRows {
			fmt.Println("not found rows")
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
func (m *DBUserRepositore) UpdateUserStatus(name string, status bool) error {
	if m == nil || m.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	query := `
	UPDATE users
	SET status=$1
	WHERE name=$2;
	`
	_, err := m.db.Exec(query, status, name)
	if err != nil {
		fmt.Println("error update status by name")
		return err
	}
	fmt.Println("status update")
	return nil
}
func (m *DBUserRepositore) ListUsers() []model.User {
	query := `
SELECT name,status
FROM users
ORDER BY name;
`
	rows, err := m.db.Query(query)
	if err != nil {
		fmt.Println("error display users")
		return nil
	}
	users := make([]model.User, 0)
	defer rows.Close()
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Name, &user.Status)
		if err != nil {
			fmt.Println("error scan")
		}
		users = append(users, user)
	}
	return users
}
func (m *DBUserRepositore) DeleteUser(name string) error {
	query := `
	DELETE FROM users
	WHERE name=$1;
	`
	_, err := m.db.Exec(query, name)
	if err != nil {
		fmt.Println("error delete")
		return err
	}
	return nil
}
