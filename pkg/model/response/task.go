package responsemodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID        primitive.ObjectID `json:"id"`
	Status    bool               `json:"status"`
	Content   string             `json:"content"`
	CreatedAt *time.Time         `json:"createdAt"`
	UpdatedAt *time.Time         `json:"updatedAt"`
}
