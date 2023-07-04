package main

import (
	"fmt"
	"net/http"
	"os"

	"ahsmha/notes/handler"
	"ahsmha/notes/injector"

	"github.com/brpaz/echozap"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := createMux()
	setupRouting(e)

	e.Validator = &handler.CustomValidator{Validator: validator.New()}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONT_URL")},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	zapLogger, _ := zap.NewProduction()
	e.Use(echozap.ZapLogger(zapLogger))
	e.Use(middleware.BodyDump(bodyDumpHandler))

	return e
}

func setupRouting(e *echo.Echo) {
	noteHandler := injector.InjectNoteHandler()
	handler.InitNoteRouting(e, noteHandler)

	authHandler := injector.InjectAuthHandler()
	handler.InitAuthRouting(e, authHandler)
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Println("Request Body:", string(reqBody))
	fmt.Println("Response Body:", string(resBody))
}
