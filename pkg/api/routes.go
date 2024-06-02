package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\ntime:${time_rfc3339}\nmethod=${method}, uri=${uri}, status=${status}, bytes_in:${bytes_in}, bytes_out:${bytes_out}\n",
	}))
	e.Use(middleware.Recover())

	e.File("/", "frontend/dist/index.html")
	e.GET("/api/hello", HelloHandler)
	e.Static("/assets", "frontend/dist/assets/")

	return e
}
