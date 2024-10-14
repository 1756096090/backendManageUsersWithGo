package repositories

import (
	"context"
	"FirstProyectWebEngineering/config"
	"FirstProyectWebEngineering/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct{}

func (r *EmployeeRepository) Create(employee *models.Employee) (*mongo.InsertOneResult, error) {
	collection := config.DB.Collection("employees")
	return collection.InsertOne(context.Background(), employee)
}

func (r *EmployeeRepository) FindAll() ([]models.Employee, error) {
	collection := config.DB.Collection("employees")
	var employees []models.Employee
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var employee models.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *EmployeeRepository) FindByID(id primitive.ObjectID) (models.Employee, error) {
	collection := config.DB.Collection("employees")
	var employee models.Employee
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&employee)
	return employee, err
}

func (r *EmployeeRepository) Update(id primitive.ObjectID, employee *models.Employee) (*mongo.UpdateResult, error) {
	collection := config.DB.Collection("employees")
	return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": employee})
}

func (r *EmployeeRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := config.DB.Collection("employees")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}

