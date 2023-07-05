package handler

import (
	"ahsmha/notes/middleware"

	"github.com/labstack/echo/v4"
)

func InitNoteRouting(e *echo.Echo, noteHandler NoteHandler) {
	e.GET("/api/notes", noteHandler.Show())              //, middleware.IsAuthenticated)
	e.POST("/api/notes", noteHandler.Create())           //, middleware.IsAuthenticated)
	e.DELETE("/api/notes/:noteID", noteHandler.Delete()) //, middleware.IsAuthenticated)
}
func InitAuthRouting(e *echo.Echo, authHandler AuthHandler) {
	e.GET("/api/auth", authHandler.Get())
	e.POST("/api/login", authHandler.Create())
	e.POST("/api/logout", authHandler.Delete(), middleware.IsAuthenticated)
	e.POST("/api/signup", authHandler.CreateUser())
}
