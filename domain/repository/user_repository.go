package repository

import "ahsmha/Tasks/domain/model"

type UserRepository interface {
	GetById(id string) (*model.User, error)
	Create(user *model.User) error
}
