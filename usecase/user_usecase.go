package usecase

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
)

type UserUsecase interface {
	GetById(id int) (user *model.User, err error)
	Create(user *model.User) error
}

// Create struct with a lowercase letter because we only refer to the repository from the use case
type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (usecase *userUsecase) GetById(id int) (user *model.User, err error) {
	user, err = usecase.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (usecase *userUsecase) Create(user *model.User) error {
	err := usecase.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}
