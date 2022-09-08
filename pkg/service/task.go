package service

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"todoApp/internal/dao"
	"todoApp/internal/model"
	modelapi "todoApp/pkg/model/api"
	responsemodel "todoApp/pkg/model/response"
)

type Task struct {
}

// CreateTask ...
func (s Task) CreateTask(ctx context.Context, payload modelapi.CreateTask) (result *responsemodel.ResponseCreate, err error) {
	var (
		d = dao.Task{}
	)
	task := payload.NewTask()

	err = d.InsertOne(ctx, task)
	if err != nil {
		return nil, errors.New("Common_Error_When_Handle")
	}
	return &responsemodel.ResponseCreate{ID: task.ID.Hex()}, nil
}

// UpdateTask ...
func (s Task) UpdateTask(ctx context.Context, task model.Task) (*responsemodel.ResponseUpdate, error) {
	var d = dao.Task{}

	if err := d.UpdateByID(ctx, task.ID, bson.M{
		"$set": bson.M{
			"status": !task.Status,
		},
	}); err != nil {
		return nil, errors.New("common_error_when_handler")
	}
	return &responsemodel.ResponseUpdate{ID: task.ID.Hex()}, nil
}

// Delete ...
func (s Task) Delete(ctx context.Context, task model.Task) (*responsemodel.ResponseDelete, error) {
	var (
		d = dao.Task{}
	)

	total := d.DeleteOne(ctx, task.ID)

	fmt.Println(total)
	return &responsemodel.ResponseDelete{ID: task.ID.Hex()}, nil
}

// All ...
func (s Task) All(ctx context.Context) (r responsemodel.ResponseList) {
	var (
		wg sync.WaitGroup
		d  = dao.Task{}
	)
	r.List = make([]responsemodel.Task, 0)
	cond := bson.D{}

	wg.Add(1)

	go func() {
		docs := d.FindByCondition(ctx, cond)
		r.List = s.getListResponse(docs)
		defer wg.Done()
	}()
	wg.Wait()

	return r
}

func (s Task) getListResponse(docs []model.Task) []responsemodel.Task {
	var (
		wg sync.WaitGroup
	)
	result := make([]responsemodel.Task, len(docs))
	wg.Add(len(docs))
	for i, value := range docs {
		go func(index int, task model.Task) {
			defer wg.Done()
			result[index] = s.getResponse(task)
		}(i, value)
	}
	wg.Wait()

	return result
}

func (s Task) getResponse(task model.Task) responsemodel.Task {
	return responsemodel.Task{
		ID:        task.ID,
		Status:    task.Status,
		Content:   task.Content,
		CreatedAt: &task.CreatedAt,
		UpdatedAt: &task.UpdatedAt,
	}
}
