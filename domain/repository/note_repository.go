package repository

import "ahsmha/notes/domain/model"

type NoteRepository interface {
	GetAllNotesByEmail(email string) (*[]model.Note, error)
	Create(note *model.Note) (int64, error)
	Update(note *model.Note) error
	Delete(id int, email string) error
}
