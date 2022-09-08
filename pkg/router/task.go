package router

import (
	"github.com/labstack/echo/v4"
	"todoApp/pkg/handler"
	"todoApp/pkg/router/checkexist"
)

func task(e *echo.Echo) {

	g := e.Group("/task")
	h := handler.Task{}

	g.GET("", h.All)

	g.POST("", h.CreateTask)

	g.PATCH("/un-do/:id", h.UnDo, checkexist.Task)

	g.DELETE("/:id", h.DeleteTask, checkexist.Task)
}
