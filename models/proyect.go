package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Proyect struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Description string `bson:"description"`
}
