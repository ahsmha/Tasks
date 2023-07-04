package infra

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
)

type NoteRepository struct {
	SqlHandler
}

// Only guaranteed to return an interface of repository.ArticleRepository
// In other words, as long as there is an interface, it is possible to mock without using sqlHandler and depending on DB
func NewNoteRepository(sqlHandler SqlHandler) repository.NoteRepository {
	return &NoteRepository{
		SqlHandler: sqlHandler,
	}
}

func (noteRepository *NoteRepository) GetById(id int) (*model.Note, error) {
	return nil, nil
}

func (noteRepository *NoteRepository) Delete(id int) error {
	return nil
}

func (noteRepository *NoteRepository) Update(note *model.Note) error {
	return nil
}

func (noteRepository *NoteRepository) Create(note *model.Note) (int64, error) {
	return 0, nil
}
