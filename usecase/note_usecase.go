package usecase

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
)

type NoteUsecase interface {
	GetById(id int) (note *model.Note, err error)
	Create(note *model.Note) (int64, error)
	Update(note *model.Note) error
	Delete(id int) error
}

// Since the repository is only referenced from the usecase, create a lowercase struct
type noteUsecase struct {
	noteRepo repository.NoteRepository
}
