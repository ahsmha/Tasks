package repository

import "ahsmha/Tasks/domain/model"

type TaskRepository interface {
	GetAllCreatedTasks(id int) (*[]model.Task, error)
	GetAllAssignedTasks(id int) (*[]model.Task, error)
	Create(Task *model.Task, Id int) error
	Update(Task *model.Task) error
	Delete(id int, email string) error
}
