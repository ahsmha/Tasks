package model

import "time"

type Note struct {
	ID      int       `db:"id" json:"id"`
	Title   string    `db:"title" validate:"required,max=50" json:"title"`
	Body    string    `db:"body" validate:"required" json:"body"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}
