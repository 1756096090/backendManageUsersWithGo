package services

import (
	"FirstProyectWebEngineering/models"
	"FirstProyectWebEngineering/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	Repository repositories.EmployeeRepository
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (*mongo.InsertOneResult, error) {
    employee.ID = primitive.NewObjectID()
    return s.Repository.Create(employee)
}

func (s *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return s.Repository.FindAll()
}

func (s *EmployeeService) GetEmployeeByID(id primitive.ObjectID) (models.Employee, error) {
	return s.Repository.FindByID(id)
}

func (s *EmployeeService) UpdateEmployee(id primitive.ObjectID, employee *models.Employee) (*mongo.UpdateResult, error) {
	return s.Repository.Update(id, employee)
}

func (s *EmployeeService) DeleteEmployee(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.Repository.Delete(id)
}
