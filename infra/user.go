package infra

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
	"database/sql"
	"time"
)

func UserCreate(user *model.User) (sql.Result, error) {
	return nil, nil
}

func UserGetByName(name string) (*model.User, error) {
	return nil, nil
}

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	return UserRepository{SqlHandler: sqlHandler}
}

func (userRepository UserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (name, password, created, updated) VALUES (?, ?, ?, ?);`

	now := time.Now()

	if _, err := userRepository.SqlHandler.Conn.Exec(query, user.Name, user.Password, now, now); err != nil {
		return err
	}

	return nil
}

func (userRepository UserRepository) GetByName(name string) (*model.User, error) {
	query := `SELECT * FROM users WHERE name = ?;`

	var user model.User
	if err := userRepository.SqlHandler.Conn.Get(&user, query, name); err != nil {
		return nil, err
	}

	return &user, nil
}
