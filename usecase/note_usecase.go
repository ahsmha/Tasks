package usecase

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
)

type NoteUsecase interface {
	GetAllNotesByUser(id int) (note *model.Note, err error)
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

func (usecase *noteUsecase) GetAllNotesByUser(id int) (note *model.Note, err error) {
	note, err = usecase.noteRepo.GetAllNotesByUser(id)
	if err != nil {
		return nil, err
	}

	return note, err
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

func (usecase *noteUsecase) Delete(id int) error {
	err := usecase.noteRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
