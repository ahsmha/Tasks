package usecase

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
)

type UserUsecase interface {
	GetByName(name string) (user *model.User, err error)
	Create(user *model.User) error
}

// Create struct with a lowercase letter because we only refer to the repository from the use case
type userUsecase struct {
	userRepo repository.UserRepository
}
