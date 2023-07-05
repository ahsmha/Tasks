package repository

import "ahsmha/notes/domain/model"

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
}
