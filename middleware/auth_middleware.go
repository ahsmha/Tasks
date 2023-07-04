package middleware

import (
	"ahsmha/notes/util"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			message := "cookie is empty"
			fmt.Println(message, err)

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": message,
			})
		}

		if err := util.ParseJwt(cookie.Value); err != nil {
			message := "cookie parse failed"
			fmt.Println(message, err)

			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": message,
			})
		}

		return next(c)
	}
}
