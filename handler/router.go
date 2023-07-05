package handler

import (
	"github.com/labstack/echo/v4"
)

func InitNoteRouting(e *echo.Echo, noteHandler NoteHandler) {
	e.GET("/api/notes/:userID", noteHandler.Show())
	e.POST("/api/notes", noteHandler.Create())
	e.DELETE("/api/notes", noteHandler.Delete())
}

func InitAuthRouting(e *echo.Echo, authHandler AuthHandler) {
	// e.GET("/api/auth", authHandler.Get())
	// e.POST("/api/login", authHandler.Create())
	// e.POST("/api/logout", authHandler.Delete(), middleware.IsAuthenticated)
	// e.POST("/api/signup", authHandler.CreateUser())
}
