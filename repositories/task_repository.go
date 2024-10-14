package repositories

import (
	"context"
	"FirstProyectWebEngineering/config"
	"FirstProyectWebEngineering/models"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct{}

func (r *TaskRepository) Create(task *models.Task) (*mongo.InsertOneResult, error) {
	collection := config.DB.Collection("tasks")
	return collection.InsertOne(context.Background(), task)
}

func (r *TaskRepository) FindAll() ([]models.Task, error) {
	collection := config.DB.Collection("tasks")
	var tasks []models.Task
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) FindByID(id primitive.ObjectID) (models.Task, error) {
	collection := config.DB.Collection("tasks")
	var task models.Task
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	return task, err
}

func (r *TaskRepository) Update(id primitive.ObjectID, task *models.Task) (*mongo.UpdateResult, error) {
	collection := config.DB.Collection("tasks")
	return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": task})
}

func (r *TaskRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := config.DB.Collection("tasks")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}

func (r *TaskRepository) FindTasksBetweenDates(startDate, endDate time.Time) ([]models.Task, error) {
	collection := config.DB.Collection("tasks")
	var tasks []models.Task

	filter := bson.M{
		"$or": []bson.M{
			{"start_date": bson.M{"$gte": startDate, "$lte": endDate}},
			{"end_date": bson.M{"$gte": startDate, "$lte": endDate}},
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
