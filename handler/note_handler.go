package handler

import (
	"ahsmha/notes/domain/model"
	"ahsmha/notes/usecase"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
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

type NoteRequest struct {
	Note      model.Note `json:"note,omitempty"`
	SessionId string     `json:"sid"`
}

type AllNotesOutput struct {
	Notes   []model.Note `json:"notes"`
	Message string       `json:"message"`
}

func (handler *NoteHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		var sid NoteRequest
		var out AllNotesOutput
		if err := c.Bind(&sid); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "request format is invalid"
			return c.JSON(http.StatusBadRequest, out)
		}
		token, err := handler.parseJwt(sid.SessionId)
		if err != nil {
			c.Logger().Error(err.Error())

			out.Message = "invalid token"
			return c.JSON(http.StatusUnprocessableEntity, out)
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["iss"].(string)

		notes, err := handler.noteUsecase.GetAllNotesByEmail(email)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusNoContent, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"notes": *notes,
		})
	}
}

func (handler *NoteHandler) parseJwt(tokenString string) (*jwt.Token, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		message := "JWT parsing failed"
		fmt.Println(message, err)

		return nil, err
	}
	if token.Claims.Valid() != nil {
		fmt.Println("Invalid JWT token:", token.Claims.Valid())
		return nil, errors.New("Invalid JWT token")
	}
	return token, nil
}

type NoteOutput struct {
	NoteId  int64  `json:"id"`
	Message string `json:"message"`
}

func (handler *NoteHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			note NoteRequest
			out  NoteOutput
		)

		if err := c.Bind(&note); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "request format is invalid"
			return c.JSON(http.StatusBadRequest, out)
		}

		token, err := handler.parseJwt(note.SessionId)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"message": "invalid token",
				"error":   err,
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		note.Note.Email = claims["iss"].(string)

		id, err := handler.noteUsecase.Create(&note.Note)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
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
		var sid NoteRequest
		if err := c.Bind(&sid); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusBadRequest, err)
		}

		token, err := handler.parseJwt(sid.SessionId)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnprocessableEntity, echo.Map{
				"message": "invalid token",
				"error":   err,
			})
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["iss"].(string)
		if err := handler.noteUsecase.Delete(id, email); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Note deleted successfully",
			"id":      id,
		})
	}
}
