package router

import (
	"github.com/labstack/echo/v4"
	"todoApp/internal/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.CORSConfig())

	task(e)
}
