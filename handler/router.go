package handler

import (
	"github.com/labstack/echo/v4"
)

func InitTaskRouting(e *echo.Echo, taskHandler TaskHandler) {
	e.GET("/api/tasks", taskHandler.Show()) //, middleware.IsAuthenticated)
	e.POST("/api/tasks", taskHandler.Create())
	// e.DELETE("/api/tasks/:taskID", taskHandler.Delete()) //, middleware.IsAuthenticated)
}

// func InitAuthRouting(e *echo.Echo, authHandler AuthHandler) {
// e.GET("/api/auth", authHandler.Get())
// e.POST("/api/login", authHandler.Create())
// e.POST("/api/logout", authHandler.Delete(), middleware.IsAuthenticated)
// e.POST("/api/signup", authHandler.CreateUser())
// }
