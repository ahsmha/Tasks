package usecase

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
)

type NoteUsecase interface {
	GetAllNotesByEmail(email string) (notes *[]model.Note, err error)
	Create(note *model.Note) (int64, error)
	Update(note *model.Note) error
	Delete(id int, email string) error
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

func (usecase *noteUsecase) GetAllNotesByEmail(email string) (notes *[]model.Note, err error) {
	notes, err = usecase.noteRepo.GetAllNotesByEmail(email)
	if err != nil {
		return nil, err
	}

	return notes, err
}
func (usecase *noteUsecase) Create(note *model.Note) (int64, error) {
	id, err := usecase.noteRepo.Create(note)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (usecase *noteUsecase) Update(note *model.Note) error {
	return nil
}

func (usecase *noteUsecase) Delete(id int, email string) error {
	err := usecase.noteRepo.Delete(id, email)
	if err != nil {
		return err
	}

	return nil
}
