package usecase

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
	"ahsmha/Tasks/utils"
	"errors"
)

type TaskUsecase interface {
	GetAllTasksByRole(role string, id int) (Tasks *[]model.Task, err error)
	Create(Task *model.Task) (int64, error)
	Update(Task *model.Task) error
	Delete(id int, email string) error
}

// Since the repository is only referenced from the usecase, create a lowercase struct
type taskUsecase struct {
	TaskRepo repository.TaskRepository
}

func NewTaskUsecase(TaskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		TaskRepo: TaskRepo,
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

func (usecase *taskUsecase) Create(Task *model.Task) (int64, error) {
	id, err := usecase.TaskRepo.Create(Task)
	if err != nil {
		return 0, err
	}

	return id, err
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
