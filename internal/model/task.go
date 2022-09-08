package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	Status    bool               `bson:"status"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
