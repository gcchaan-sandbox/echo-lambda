package app

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func Router(e *echo.Echo, req events.APIGatewayProxyRequest) {
	e.GET("/api/posts/:channel", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})
	e.GET("/api/posts", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})
	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})
}
