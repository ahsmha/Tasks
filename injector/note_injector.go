package injector

import (
	"ahsmha/notes/domain/repository"
	"ahsmha/notes/handler"
	"ahsmha/notes/infra"
	"ahsmha/notes/usecase"
)

func InjectNoteHandler() handler.NoteHandler {
	return handler.NewNoteHandler(InjectNoteUsecase())
}

func InjectNoteUsecase() usecase.NoteUsecase {
	noteRepository := InjectNoteRepository()

	return usecase.NewNoteUsecase(noteRepository)
}

func InjectNoteRepository() repository.NoteRepository {
	sqlHandler := InjectDB()

	return infra.NewNoteRepository(sqlHandler)
}
