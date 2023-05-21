package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	WorkerName string `bson:"workerName,omitempty"`
	Date string `bson:"date,omitempty"`
	Hours int `bson:"hours,omitempty"`
}