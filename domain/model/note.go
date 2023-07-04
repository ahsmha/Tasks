package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Note struct {
	ID      uint32    `db:"id" json:"id"`
	Title   string    `db:"title" validate:"required,max=50" json:"title"`
	Body    string    `db:"body" validate:"required" json:"body"`
	UserId  uint32    `db:"user_id" json:"uid"`
	Created time.Time `db:"created" json:"created"`
	Updated time.Time `db:"updated" json:"updated"`
}

func (a *Note) ValidationErrors(err error) []string {
	var errMessages []string

	for _, err := range err.(validator.ValidationErrors) {
		var message string

		// Switching validation messages for each tag
		switch err.Field() {
		case "Title":
			switch err.Tag() {
			case "required":
				message = "Title is required"
			case "max":
				message = "Title can be upto 50 chars"
			}
		case "Body":
			message = "Body is required"
		}

		if message != "" {
			errMessages = append(errMessages, message)
		}
	}

	return errMessages
}
