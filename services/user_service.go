package services

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
    Repository repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.User) (*mongo.InsertOneResult, error) {
    user.ID = primitive.NewObjectID()
    return s.Repository.Create(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repository.FindAll()
}

func (s *UserService) GetUserByID(id primitive.ObjectID) (models.User, error) {
    return s.Repository.FindByID(id)
}

func (s *UserService) UpdateUser(id primitive.ObjectID, user *models.User) (*mongo.UpdateResult, error) {
    return s.Repository.Update(id, user)
}

func (s *UserService) DeleteUser(id primitive.ObjectID) (*mongo.DeleteResult, error) {
    return s.Repository.Delete(id)
}
