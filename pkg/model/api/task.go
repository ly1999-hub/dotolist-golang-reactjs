package apimodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todoApp/internal/model"
)

type CreateTask struct {
	Content string `json:"content" xml:"content" form:"content" query:"content"`
}

func (c CreateTask) NewTask() model.Task {
	return model.Task{
		ID:        primitive.NewObjectID(),
		Status:    false,
		Content:   c.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
