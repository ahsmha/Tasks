package repository

import "ahsmha/Tasks/domain/model"

type UserRepository interface {
	GetById(id int) (*model.User, error)
	Create(user *model.User) error
}
