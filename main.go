package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ahsmha/Tasks/handler"
	"ahsmha/Tasks/injector"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e := createMux()
	setupRouting(e)

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
	e.Use(middleware.BodyDump(bodyDumpHandler))

	return e
}

func setupRouting(e *echo.Echo) {
	noteHandler := injector.InjectTaskHandler()
	handler.InitTaskRouting(e, noteHandler)

	// authHandler := injector.InjectAuthHandler()
	// handler.InitAuthRouting(e, authHandler)
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Println("Request Body:", string(reqBody))
	fmt.Println("Response Body:", string(resBody))
}
