package usecase

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
	"ahsmha/Tasks/utils"
	"errors"
)

type TaskUsecase interface {
	GetAllTasksByRole(role string, id int) (Tasks *[]model.Task, err error)
	Create(Task *model.Task) error
	Update(Task *model.Task) error
	Delete(id int, email string) error
}

// Since the repository is only referenced from the usecase, create a lowercase struct
type taskUsecase struct {
	TaskRepo repository.TaskRepository
	UserRepo repository.UserRepository
}

func NewTaskUsecase(TaskRepo repository.TaskRepository, UserRepo repository.UserRepository) TaskUsecase {
	return &taskUsecase{
		TaskRepo: TaskRepo,
		UserRepo: UserRepo,
	}
}

func (usecase *taskUsecase) GetAllTasksByRole(role string, id int) (Tasks *[]model.Task, err error) {
	if role == utils.LEAD_ROLE {
		Tasks, err := usecase.TaskRepo.GetAllAssignedTasks(id)
		if err != nil {
			return nil, err
		}
		return Tasks, nil
	} else if role == utils.SUBORDINATE_ROLE {
		Tasks, err := usecase.TaskRepo.GetAllCreatedTasks(id)
		if err != nil {
			return nil, err
		}
		return Tasks, nil
	}
	return nil, errors.New("undefined role")
}

func (usecase *taskUsecase) Create(Task *model.Task) error {
	// check if role = lead for this user_id
	user, err := usecase.UserRepo.GetById(Task.Creator_id)
	if err != nil {
		return err
	}
	if user.Role != utils.LEAD_ROLE {
		return errors.New("user is not a lead")
	}
	return usecase.TaskRepo.Create(Task, Task.Creator_id)
}

func (usecase *taskUsecase) Update(Task *model.Task) error {
	err := usecase.TaskRepo.Update(Task)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *taskUsecase) Delete(id int, email string) error {
	err := usecase.TaskRepo.Delete(id, email)
	if err != nil {
		return err
	}

	return nil
}
