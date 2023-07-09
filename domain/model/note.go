package model

import (
	"time"
)

type Note struct {
	ID      uint32    `json:"id"`
	Title   string    `validate:"required,max=50" json:"title"`
	Body    string    `validate:"required" json:"body"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
