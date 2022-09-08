package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todoApp/internal/logger"
	"todoApp/internal/model"
	"todoApp/internal/module"
)

type Task struct {
}

func (d Task) InsertOne(ctx context.Context, task model.Task) error {
	var (
		col = d.getCollection()
	)
	_, err := col.InsertOne(ctx, task)
	if err != nil {
		logger.Error("Error InsertOne-Task", logger.LogData{
			"task":  task,
			"error": err.Error(),
		})
	}
	return err
}
func (d Task) FindOne(ctx context.Context, cond interface{}) (doc model.Task) {
	var col = d.getCollection()

	err := col.FindOne(ctx, cond).Decode(&doc)
	if err != nil {
		logger.Error("error FindOne-Task", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return doc
}

func (d Task) FindByID(ctx context.Context, id primitive.ObjectID) (doc model.Task) {
	return d.FindOne(ctx, bson.M{"_id": id})
}
func (d Task) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	var col = d.getCollection()

	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("Error UpdateOne-Task", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"error":   err.Error(),
		})
	}
	return err
}

func (d Task) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, payload)
}

func (d Task) DeleteOne(ctx context.Context, id primitive.ObjectID) int64 {
	var col = d.getCollection()
	res, err := col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		logger.Error("Error DeleteOne-Task", logger.LogData{
			"id":  id,
			"err": err.Error(),
		})
	}
	return res.DeletedCount
}

func (d Task) FindByCondition(ctx context.Context, cond interface{}) (docs []model.Task) {
	col := d.getCollection()
	cursor, err := col.Find(ctx, cond)
	if err != nil {
		logger.Error("Error FindByCondition-Task", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		logger.Error("dao.Task - FindByCondition - decode", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
		return nil
	}
	return docs
}

func (d Task) getCollection() *mongo.Collection {
	db := module.GetInstance()
	return db.Collection("todolist")
}
