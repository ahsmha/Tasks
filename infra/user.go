package infra

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
	"database/sql"
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
	return nil
}

func (userRepository UserRepository) GetByName(name string) (*model.User, error) {
	return nil, nil
}
