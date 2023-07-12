package repository

import "ahsmha/Tasks/domain/model"

type TaskRepository interface {
	GetAllCreatedTasks(id int) (*[]model.TaskExtended, error)
	GetAllAssignedTasks(id int) (*[]model.TaskExtended, error)
	GetById(id int) (*model.TaskExtended, error)
	Create(Task *model.Task, Id int) error
	Delete(id int, email string) error
	UpdateTaskByLead(Task *model.Task, Id int) error
	UpdateStatusBySubOrdinate(Status string, Id int) error
}
