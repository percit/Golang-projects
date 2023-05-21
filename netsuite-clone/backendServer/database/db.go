package database
import (
	"log"
	"context"

	"backendServer/database/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const dbName = "netsuite"
const collectionName = "users"
var mongoClient *mongo.Client

type DB struct {
	collection *mongo.Collection //grouping of mongodb documents(mongodb stores data as documents ->bson documents)
}

func (db *DB) InitDB(connectionString string) {

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongoDB connection success")
	
	db.collection = client.Database(dbName).Collection(collectionName)
	log.Println("collection instance is ready")


	mongoClient = client //this is for closing connection
}

func (db *DB) CloseConnection() {
	err := mongoClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}

func (db *DB) InsertUser(model model.User) {
	insertedId, err := db.collection.InsertOne(context.Background(), model)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted 1 model with id", insertedId.InsertedID)
}

func (db *DB) UpdateOneRecord(hours int, userID primitive.ObjectID) {
	// id, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"hours":hours}}//TODO check if this works

	result, err := db.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("modified count", result.ModifiedCount)
}

func (db *DB) FindUser(userName string) model.User {
	filter := bson.M{"workerName": userName}
	var result model.User
	err := db.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (db *DB) GetAllRecords() []*model.User {
	cursor, err := db.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err) 
	}
	var users []*model.User
	for cursor.Next(context.Background()){
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	defer cursor.Close(context.Background())
	return users
}

func (db *DB) DeleteOneRecord(userID string) { //primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(userID) 
	filter := bson.M{"_id": id}
	deleteCount, err := db.collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("modified count", deleteCount.DeletedCount)
}

func (db *DB) DeleteAllRecords() {
	deleteCount, err := db.collection.DeleteMany(context.Background(), bson.M{},nil)

	if err != nil {
		log.Fatal(err) 
	}
	log.Println("modified count", deleteCount.DeletedCount) 
}
