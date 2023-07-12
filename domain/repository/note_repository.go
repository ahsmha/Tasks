package repository

import "ahsmha/Tasks/domain/model"

type TaskRepository interface {
	GetAllTasksByRole(role string) (*[]model.Task, error)
	Create(Task *model.Task) (int64, error)
	Update(Task *model.Task) error
	Delete(id int, email string) error
}
