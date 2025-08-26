package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (email, password_hash, name) VALUES ($1, $2, $3) RETURNING id`

	return m.DB.QueryRowContext(ctx, query, user.Email, user.PasswordHash, user.Name).Scan(&user.Id)

}

func (m *UserModel) getUser(query string, args ...interface{}) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// query := `SELECT id, email, name, password_hash FROM users WHERE id = $1`

	var user User

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func (m *UserModel) Get(id int) (*User, error) {
	query := `SELECT id, email, name, password_hash FROM users WHERE id = $1`
	return m.getUser(query, id)
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := `SELECT id, email, name, password_hash FROM users WHERE email = $1`
	return m.getUser(query, email)
}
