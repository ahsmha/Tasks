package infra

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/domain/repository"
	"fmt"
	"time"
)

type TaskRepository struct {
	SqlHandler
}

func NewTaskRepository(sqlHandler SqlHandler) repository.TaskRepository {
	return &TaskRepository{
		SqlHandler: sqlHandler,
	}
}

func (TaskRepository *TaskRepository) GetAllCreatedTasks(id int) (*[]model.Task, error) {
	var Tasks []model.Task
	TaskQuery := `
		SELECT id, title, due_date, status
		FROM tasks
		WHERE creator_id = ?;
	`

	rows, err := TaskRepository.SqlHandler.Conn.Queryx(TaskQuery, id)
	if err != nil {
		err = fmt.Errorf("failed to select Tasks: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Task model.Task
		err := rows.StructScan(&Task)
		if err != nil {
			err = fmt.Errorf("failed to scan Task: %w", err)
			return nil, err
		}
		Tasks = append(Tasks, Task)
	}

	if err := rows.Err(); err != nil {
		err = fmt.Errorf("failed to iterate over Tasks: %w", err)
		return nil, err
	}

	return &Tasks, nil

}

func (TaskRepository *TaskRepository) GetAllAssignedTasks(id int) (*[]model.Task, error) {
	var Tasks []model.Task
	TaskQuery := `
		SELECT t.id, t.title, t.due_date, t.status
		FROM tasks t
		JOIN task_assignee_mapping tam
		ON t.id = tam.task_id
		WHERE tam.assignee_id = ?;
	`

	rows, err := TaskRepository.SqlHandler.Conn.Queryx(TaskQuery, id)
	if err != nil {
		err = fmt.Errorf("failed to select Tasks: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Task model.Task
		err := rows.StructScan(&Task)
		if err != nil {
			err = fmt.Errorf("failed to scan Task: %w", err)
			return nil, err
		}
		Tasks = append(Tasks, Task)
	}

	if err := rows.Err(); err != nil {
		err = fmt.Errorf("failed to iterate over Tasks: %w", err)
		return nil, err
	}

	return &Tasks, nil

}

func (TaskRepository *TaskRepository) Delete(id int, email string) error {
	query := `DELETE FROM Tasks WHERE id = ? AND email = ?;`

	tx := TaskRepository.SqlHandler.Conn.MustBegin()

	_, err := tx.Exec(query, id, email)
	if err != nil {
		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (TaskRepository *TaskRepository) Create(Task *model.Task, Id int) error {
	now := time.Now()

	Task.Created = now
	Task.Updated = now

	query := `INSERT INTO Tasks (title, due_date, status, creator_id, created, updated)
	VALUES (:title, :due_date, :status, :creator_id, :created, :updated);`

	// start transaction
	tx := TaskRepository.SqlHandler.Conn.MustBegin()
	_, err := tx.NamedExec(query, Task)
	if err != nil {
		err = fmt.Errorf("failed to create Task: %w", err)
		_ = tx.Rollback()

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (TaskRepository *TaskRepository) UpdateTaskByLead(Task *model.Task, Id int) error {
	now := time.Now()
	Task.Updated = now
	Task.ID = Id

	query := `UPDATE tasks
	SET title = :title,
		due_date = :due_date,
		updated = :updated,
		status = :status,
	WHERE id = id;`

	tx := TaskRepository.SqlHandler.Conn.MustBegin()

	_, err := tx.NamedExec(query, Task)
	if err != nil {
		err = fmt.Errorf("failed to update tasks: %w", err)
		_ = tx.Rollback()

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (TaskRepository *TaskRepository) UpdateStatusBySubOrdinate(Status string, Id int) error {

	query := `UPDATE tasks
	SET status = ?,
	WHERE id = ?;
	`
	tx := TaskRepository.SqlHandler.Conn.MustBegin()

	_, err := tx.Exec(query, Status, Id)
	if err != nil {
		err = fmt.Errorf("failed to update tasks: %w", err)
		_ = tx.Rollback()

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
