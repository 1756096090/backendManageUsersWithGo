package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Age  int `bson:"age"`
	Email string `bson:"email"`
    Password string `bson:"password"`
}

