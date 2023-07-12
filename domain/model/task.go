package model

import (
	"time"
)

type Task struct {
	ID         int       `json:"id"`
	Title      string    `validate:"required,max=50" json:"title"`
	Due_date   string    `json:"due_date"`
	Status     string    `json:"status"`
	Creator_id int       `json:"creator_id"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

type TaskExtended struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `validate:"required,max=50" json:"title,omitempty"`
	Due_date    string    `json:"due_date,omitempty"`
	Status      string    `json:"status,omitempty"`
	Creator_id  int       `json:"creator_id,omitempty"`
	Assignee_id int       `json:"assignee_id,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}
