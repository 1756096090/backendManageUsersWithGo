package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	ID_Project  primitive.ObjectID `bson:"id_project"`
	ID_Employee primitive.ObjectID `bson:"id_employee"`
	StartDate   time.Time          `bson:"start_date"`
	EndDate     time.Time          `bson:"end_date"`
	IsCompleted bool               `bson:"is_completed"`
}
