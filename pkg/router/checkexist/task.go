package checkexist

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todoApp/internal/dao"
	"todoApp/internal/response"
)

func Task(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx = c.Request().Context()
			id  = c.Param("id")
			d   = dao.Task{}
		)

		fmt.Println(id)
		idTask, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return response.R400(c, nil, err.Error())
		}
		task := d.FindByID(ctx, idTask)

		if task.ID.IsZero() {
			return response.R400(c, nil, err.Error())
		}

		c.Set("task", task)

		return next(c)
	}
}
