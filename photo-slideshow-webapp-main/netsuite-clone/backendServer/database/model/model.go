package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Hours struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	WorkerName string `json: "workerName, omitempty"` //rething this, maybe add new argument "worker" for every function
	Date string `json: "date, omitempty"`
	Hours int `json: "hours, omitempty"`
}