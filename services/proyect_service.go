package services

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProyectService struct {
	Repository repositories.ProyectRepository
}

func (s *ProyectService) CreateProyect(proyect *models.Proyect) (*mongo.InsertOneResult, error) {
    proyect.ID = primitive.NewObjectID()
    return s.Repository.Create(proyect)
}

func (s *ProyectService) GetAllProyects() ([]models.Proyect, error) {
	return s.Repository.FindAll()
}

func (s *ProyectService) GetProyectByID(id primitive.ObjectID) (models.Proyect, error) {
	return s.Repository.FindByID(id)
}

func (s *ProyectService) UpdateProyect(id primitive.ObjectID, proyect *models.Proyect) (*mongo.UpdateResult, error) {
	return s.Repository.Update(id, proyect)
}

func (s *ProyectService) DeleteProyect(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.Repository.Delete(id)
}
