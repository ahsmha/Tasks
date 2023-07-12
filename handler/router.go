package handler

import (
	"github.com/labstack/echo/v4"
)

func InitTaskRouting(e *echo.Echo, taskHandler TaskHandler) {
	e.GET("/api/tasks", taskHandler.Show())
	e.POST("/api/tasks", taskHandler.Create())
	e.POST("/api/tasks/edit/:taskId", taskHandler.Update())
}
