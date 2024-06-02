package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from index")
}

func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from api")
}
