package infra

import (
	"ahsmha/notes/domain/model"
	"database/sql"
)

type UserRepository struct {
	SqlHandler
}

func UserCreate(user *model.User) (sql.Result, error) {
	return nil, nil
}

func UserGetByName(name string) (*model.User, error) {
	return nil, nil
}

// more functions to be added
