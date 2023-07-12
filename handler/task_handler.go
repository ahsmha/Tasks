package handler

import (
	"ahsmha/Tasks/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type TaskHandler struct {
	TaskUsecase usecase.TaskUsecase
}

func NewTaskHandler(TaskUsecase usecase.TaskUsecase) TaskHandler {
	return TaskHandler{
		TaskUsecase: TaskUsecase,
	}
}

type TaskRequest struct {
	Role   string `json:"role"`
	UserId int    `json:"user_id"`
}

type AllTasksOutput struct {
	Message string `json:"message"`
}

func (handler *TaskHandler) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			treq TaskRequest
			out  AllTasksOutput
		)
		// validate the request
		validate := validator.New()
		if err := validate.Struct(treq); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "Invalid request body"
			return c.JSON(http.StatusBadRequest, out)
		}

		if err := c.Bind(&treq); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "could not bind request"
			return c.JSON(http.StatusBadRequest, out)
		}

		Tasks, err := handler.TaskUsecase.GetAllTasksByRole(treq.Role, treq.UserId)
		if err != nil {
			c.Logger().Error(err.Error())

			out.Message = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, out)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Tasks": *Tasks,
		})
	}
}

type TaskOutput struct {
	TaskId  int64  `json:"id"`
	Message string `json:"message"`
}

func (handler *TaskHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			Task TaskRequest
			out  TaskOutput
		)

		if err := c.Bind(&Task); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "request format is invalid"
			return c.JSON(http.StatusBadRequest, out)
		}

		// id, err := handler.TaskUsecase.Create(&Task.Task)
		// if err != nil {
		// 	c.Logger().Error(err.Error())

		// 	return c.JSON(http.StatusInternalServerError, err)
		// }

		// out.TaskId = id

		return c.JSON(http.StatusOK, out)
	}
}

func (handler *TaskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// get Task id from path parameter
		// Since it is obtained as a string type, cast it to a numeric type using the strconv package
		id, _ := strconv.Atoi(c.Param("TaskID"))
		var sid TaskRequest
		if err := c.Bind(&sid); err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "Task deleted successfully",
			"id":      id,
		})
	}
}
