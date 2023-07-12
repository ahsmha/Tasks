package handler

import (
	"ahsmha/Tasks/domain/model"
	"ahsmha/Tasks/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	Task model.Task `json:"Task,omitempty"`
	Role string     `json:"Role"`
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
		if err := c.Bind(&treq); err != nil {
			c.Logger().Error(err.Error())

			out.Message = "request format is invalid"
			return c.JSON(http.StatusBadRequest, out)
		}
		Tasks, err := handler.TaskUsecase.GetAllTasksByRole(treq.Role)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusNoContent, err)
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

		id, err := handler.TaskUsecase.Create(&Task.Task)
		if err != nil {
			c.Logger().Error(err.Error())

			return c.JSON(http.StatusInternalServerError, err)
		}

		out.TaskId = id

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
