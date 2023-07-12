package model

import (
	"time"
)

type Task struct {
	ID         int       `json:"id"`
	Title      string    `validate:"required,max=50" json:"title"`
	Due_date   time.Time `json:"due_date"`
	Status     string    `json:"status"`
	Creator_id int       `json:"creator_id"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
