package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Note struct {
	ID      uint32    `json:"id"`
	Title   string    `validate:"required,max=50" json:"title"`
	Body    string    `validate:"required" json:"body"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
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
