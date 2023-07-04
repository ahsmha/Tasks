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

func NewNoteUsecase(noteRepo repository.NoteRepository) NoteUsecase {
	return &noteUsecase{
		noteRepo: noteRepo,
	}
}

func (usecase *noteUsecase) GetById(id int) (*model.Note, error) {
	return nil, nil
}

func (usecase *noteUsecase) Create(note *model.Note) (int64, error) {
	return 0, nil
}

func (usecase *noteUsecase) Update(note *model.Note) error {
	return nil
}

func (usecase *noteUsecase) Delete(id int) error {
	return nil
}
