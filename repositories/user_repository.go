package repositories

import (
	"context"
	// "log"
	"FirstProyectWebEngineering/config"
	"FirstProyectWebEngineering/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *models.User) (*mongo.InsertOneResult, error) {
    collection := config.DB.Collection("users")
    return collection.InsertOne(context.Background(), user)
}

func (r *UserRepository) FindAll() ([]models.User, error) {
    collection := config.DB.Collection("users")
    var users []models.User
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
    for cursor.Next(context.Background()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) FindByID(id primitive.ObjectID) (models.User, error) {
    collection := config.DB.Collection("users")
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
    return user, err
}

func (r *UserRepository) Update(id primitive.ObjectID, user *models.User) (*mongo.UpdateResult, error) {
    collection := config.DB.Collection("users")
    return collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": user})
}

func (r *UserRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
    collection := config.DB.Collection("users")
    return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
