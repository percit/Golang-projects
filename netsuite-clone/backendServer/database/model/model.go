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

// const salesOnApril4th = db.getCollection('sales').find({
// 	date: { $gte: new Date('2014-04-04'), $lt: new Date('2014-04-05') }
//   }).count();
//there are dates in golang aka: new Date