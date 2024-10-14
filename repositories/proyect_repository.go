package repositories

import (
	"context"
	"FirstProyectWebEngineering/config"
	"FirstProyectWebEngineering/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProyectRepository struct{}

func (r *ProyectRepository) Create(proyect *models.Proyect) (*mongo.InsertOneResult, error) {
	collection := config.DB.Collection("proyects")
	return collection.InsertOne(context.Background(), proyect)
}

func (r *ProyectRepository) FindAll() ([]models.Proyect, error) {
	collection := config.DB.Collection("proyects")
	var proyects []models.Proyect
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var proyect models.Proyect
		cursor.Decode(&proyect)
		proyects = append(proyects, proyect)
	}
	return proyects, nil
}

func (r *ProyectRepository) FindByID(id primitive.ObjectID) (models.Proyect, error) {
	collection := config.DB.Collection("proyects")
	var proyect models.Proyect
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&proyect)
	return proyect, err
}

func (r *ProyectRepository) Update(id primitive.ObjectID, proyect *models.Proyect) (*mongo.UpdateResult, error) {
	collection := config.DB.Collection("proyects")
	return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": proyect})
}

func (r *ProyectRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := config.DB.Collection("proyects")
	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}

