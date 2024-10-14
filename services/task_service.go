package services

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TaskService struct {
	Repository repositories.TaskRepository
}

func (s *TaskService) CreateTask(task *models.Task) (*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	return s.Repository.Create(task)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repository.FindAll()
}

func (s *TaskService) GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	return s.Repository.FindByID(id)
}

func (s *TaskService) UpdateTask(id primitive.ObjectID, task *models.Task) (*mongo.UpdateResult, error) {
	return s.Repository.Update(id, task)
}

func (s *TaskService) DeleteTask(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.Repository.Delete(id)
}

func (s *TaskService) GetTasksBetweenDates(startDate, endDate time.Time) ([]models.Task, error) {
	return s.Repository.FindTasksBetweenDates(startDate, endDate)
}
