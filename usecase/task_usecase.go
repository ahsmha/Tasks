package usecase

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
)

type TaskUsecase interface {
	GetAllTasksByRole(email string) (Tasks *[]model.Task, err error)
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

func (usecase *taskUsecase) GetAllTasksByRole(role string) (Tasks *[]model.Task, err error) {
	Tasks, err = usecase.TaskRepo.GetAllTasksByRole(role)
	if err != nil {
		return nil, err
	}

	return Tasks, err
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
