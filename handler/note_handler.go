package handler

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NoteHandler struct {
	noteUsecase usecase.NoteUsecase
}

func NewNoteHandler(noteUsecase usecase.NoteUsecase) NoteHandler {
	return NoteHandler{
		noteUsecase: noteUsecase,
	}
}

func (handler *NoteHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("noteID"))
		note, err := handler.noteUsecase.GetAllNotesByUser(id)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusNoContent, err)
		}

		return c.JSON(http.StatusOK, note)
	}
}

type NoteOutput struct {
	NoteId           int64    `json:"id"`
	Message          string   `json:"message"`
	ValidationErrors []string `json:"errors"`
}

func (handler *NoteHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			note model.Note
			out  NoteOutput
		)

		if err := c.Bind(&note); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "request format is invalid"
			return c.JSON(http.StatusBadRequest, out)
		}

		if err := c.Validate(&note); err != nil {
			c.Logger().Error(err.Error())

			out.ValidationErrors = note.ValidationErrors(err)

			out.Message = "request validation failed"
			return c.JSON(http.StatusUnprocessableEntity, out)
		}

		id, err := handler.noteUsecase.Create(&note)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, out)
		}

		out.NoteId = id

		return c.JSON(http.StatusOK, out)
	}
}

func (handler *NoteHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get note id from path parameter
		// Since it is obtained as a string type, cast it to a numeric type using the strconv package
		id, _ := strconv.Atoi(c.Param("noteID"))

		if err := handler.noteUsecase.Delete(id); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, "")
		}

		return c.JSON(http.StatusOK, fmt.Sprintf("Note %d is deleted.", id))
	}
}
