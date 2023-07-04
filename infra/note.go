package infra

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/domain/repository"
	"fmt"
	"time"
)

type NoteRepository struct {
	SqlHandler
}

// Only guaranteed to return an interface of repository.NoteRepository
// In other words, as long as there is an interface, it is possible to mock without using sqlHandler and depending on DB

func NewNoteRepository(sqlHandler SqlHandler) repository.NoteRepository {
	return &NoteRepository{
		SqlHandler: sqlHandler,
	}
}

func (noteRepository *NoteRepository) GetAllNotesByUser(id int) (*model.Note, error) {
	var note model.Note
	noteQuery := `SELECT *
	FROM notes
	WHERE uid= ?;`
	if err := noteRepository.SqlHandler.Conn.Get(&note, noteQuery, id); err != nil {
		err = fmt.Errorf("failed to select note: %w", err)
		return nil, err
	}

	return &note, nil
}

func (noteRepository *NoteRepository) Delete(id int) error {
	query := `DELETE FROM notes WHERE id = ?;`

	tx := noteRepository.SqlHandler.Conn.MustBegin()

	_, err := tx.Exec(query, id)
	if err != nil {
		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (noteRepository *NoteRepository) Update(note *model.Note) error {
	return nil
}

func (noteRepository *NoteRepository) Create(note *model.Note) (int64, error) {
	now := time.Now()

	note.Created = now
	note.Updated = now

	query := `INSERT INTO notes (title, body, created, updated)
	VALUES (:title, :body, :created, :updated);`

	// start transaction
	tx := noteRepository.SqlHandler.Conn.MustBegin()
	res, err := tx.NamedExec(query, note)
	if err != nil {
		err = fmt.Errorf("failed to create note: %w", err)
		_ = tx.Rollback()

		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("failed to get last insert id: %w", err)
		return 0, err
	}

	return id, nil
}
