package database
import (
	"log"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const connectionString = ""
const dbName = "netsuite"
const collectionName = "hours"

//most important
var collection *mongo.Collection //grouping of mongodb documents(mongodb stores data as documents ->bson documents)



//########################################################################################
//########################################################################################
//########################################################################################
type DB struct {
	database map[string]int//database that will take date as string, and number of hours
}

func (db *DB) InitDB() {
	clientOption := options.Client().ApplyURI(connectionString)
	//if we connect, we always have to pass context 
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongoDB connection success")
	collection = client.Database(dbName).Collection(collectionName)
	//collection instance
	log.Println("collection instance is ready")




	db.database = make(map[string]int)
	db.database["24.04.2023"] = 8
	db.database["25.04.2023"] = 8
	db.database["26.04.2023"] = 8
	db.database["27.04.2023"] = 8
}

func (db *DB) GetHoursByDate(date string) int {
	log.Println("getHoursByDate")
	return db.database[date]
}

func (db *DB) SetHoursByDate(date string, time int) {
	log.Println("setHoursByDate")
	db.database[date] = time
}

func (db *DB) DeleteDate(date string) {
	log.Println("deleteDate")
	delete(db.database, date)
}