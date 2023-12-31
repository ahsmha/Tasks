package usecase

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
	"ahsmha/Tasks/utils"
	"database/sql"
	"errors"
	"fmt"
)

type TaskUsecase interface {
	GetAllTasksByRole(role string, id int) (Tasks *[]model.TaskExtended, err error)
	Create(Task *model.Task) error
	Update(Task *model.Task, Id int) error
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

func (usecase *taskUsecase) GetAllTasksByRole(role string, id int) (Tasks *[]model.TaskExtended, err error) {
	if role == utils.SUBORDINATE_ROLE {
		Tasks, err := usecase.TaskRepo.GetAllAssignedTasks(id)
		if err != nil {
			return nil, err
		}
		return Tasks, nil
	} else if role == utils.LEAD_ROLE {
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

func (usecase *taskUsecase) Update(Task *model.Task, Id int) error {
	// check if role is leader or subord.
	user, err := usecase.UserRepo.GetById(Task.Creator_id)
	if err != nil {
		return err
	}
	task, err := usecase.TaskRepo.GetById(Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if user.Role == utils.LEAD_ROLE {
		if task.Creator_id != user.Id {
			return fmt.Errorf("Failed because Task is not created by same user.")
		}
		err = usecase.TaskRepo.UpdateTaskByLead(Task, Id)
		if err != nil {
			return err
		}
		return nil
	} else if user.Role == utils.SUBORDINATE_ROLE {
		if err == sql.ErrNoRows {
			return fmt.Errorf("failed becasue task is not assigned to anyone")
		} else if task.Assignee_id != user.Id {
			return fmt.Errorf("cannot update because task is not assigned to user")
		}
		newStatus := Task.Status
		err := usecase.TaskRepo.UpdateStatusBySubOrdinate(newStatus, Id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("No such user found")
}

func (usecase *taskUsecase) Delete(id int, email string) error {
	err := usecase.TaskRepo.Delete(id, email)
	if err != nil {
		return err
	}

	return nil
}
