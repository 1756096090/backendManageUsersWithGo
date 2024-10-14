package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	DateOfBirth time.Time         `bson:"dateOfBirth"` 
	Email      string             `bson:"email"`
	Password   string             `bson:"password"`
}
