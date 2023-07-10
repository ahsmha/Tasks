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

func NewNoteRepository(sqlHandler SqlHandler) repository.NoteRepository {
	return &NoteRepository{
		SqlHandler: sqlHandler,
	}
}

func (noteRepository *NoteRepository) GetAllNotesByEmail(email string) (*[]model.Note, error) {
	var notes []model.Note
	noteQuery := `
		SELECT n.id,n.title,n.body,n.email,n.created,n.updated
		FROM notes n
		JOIN users u ON n.email = u.email
		WHERE n.email = ?;
	`

	rows, err := noteRepository.SqlHandler.Conn.Queryx(noteQuery, email)
	if err != nil {
		err = fmt.Errorf("failed to select notes: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note model.Note
		err := rows.StructScan(&note)
		if err != nil {
			err = fmt.Errorf("failed to scan note: %w", err)
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		err = fmt.Errorf("failed to iterate over notes: %w", err)
		return nil, err
	}

	return &notes, nil
}

func (noteRepository *NoteRepository) Delete(id int, email string) error {
	query := `DELETE FROM notes WHERE id = ? AND email = ?;`

	tx := noteRepository.SqlHandler.Conn.MustBegin()

	_, err := tx.Exec(query, id, email)
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

	query := `INSERT INTO notes (title, body, email, created, updated)
	VALUES (:title, :body, :email, :created, :updated);`

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
