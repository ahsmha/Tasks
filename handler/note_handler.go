package handler

import (
	"ahsmha/notes/usecase"
)

type NoteHandler struct {
	noteUsecase usecase.NoteUsecase
}

func NewNoteHandler(noteUsecase usecase.NoteUsecase) NoteHandler {
	return NoteHandler{
		noteUsecase: noteUsecase,
	}
}
