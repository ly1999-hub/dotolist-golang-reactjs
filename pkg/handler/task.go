package handler

import (
	"github.com/labstack/echo/v4"
	"todoApp/internal/model"
	"todoApp/internal/response"
	modelapi "todoApp/pkg/model/api"
	"todoApp/pkg/service"
)

type Task struct {
}

// CreateTask ...
func (h Task) CreateTask(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = modelapi.CreateTask{}
		s       = service.Task{}
	)
	err := c.Bind(&payload)

	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	res, errs := s.CreateTask(ctx, payload)
	if errs != nil {
		return response.R400(c, nil, errs.Error())
	}
	return response.R200(c, res, "")
}

// UnDo ...
func (h Task) UnDo(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		s    = service.Task{}
		task = c.Get("task").(model.Task)
	)

	res, err := s.UpdateTask(ctx, task)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

// DeleteTask ...
func (h Task) DeleteTask(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		s    = service.Task{}
		task = c.Get("task").(model.Task)
	)
	res, err := s.Delete(ctx, task)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res, "")
}

// All ...
func (h Task) All(c echo.Context) error {
	var (
		s   = service.Task{}
		ctx = c.Request().Context()
	)

	res := s.All(ctx)

	return response.R200(c, res, "")
}
