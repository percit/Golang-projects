package database
import (
	"log"
	"context"
	// "net/http"
	// "encoding/json"

	"backendServer/database/model"

	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const dbName = "netsuite"
const collectionName = "hours"

func (db *DB) InitDB(connectionString string) {

	clientOption := options.Client().ApplyURI(connectionString)
	//if we connect, we always have to pass context 
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mongoDB connection success")
	db.collection = client.Database(dbName).Collection(collectionName)
	//collection instance
	log.Println("collection instance is ready")




	db.database = make(map[string]int)
	db.database["24.04.2023"] = 8
	db.database["25.04.2023"] = 8
	db.database["26.04.2023"] = 8
	db.database["27.04.2023"] = 8
}

type DB struct {
	database map[string]int//database that will take date as string, and number of hours
	collection *mongo.Collection //grouping of mongodb documents(mongodb stores data as documents ->bson documents)
}
 //these are temporary for containers instead of database
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

//########################################################################################
//########################################################################################

func (db *DB) InsertRecordHelper(model model.Hours) {
	insertedId, err := db.collection.InsertOne(context.Background(), model)//ten background, to taki inny context
	if err != nil {
		log.Fatal(err)
	}
	log.Println("inserted 1 model with id", insertedId.InsertedID)
}

//########################################################################################
//########################################################################################

func (db *DB) GetAllRecords() []primitive.M{
	cursor, err := db.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err) 
	}
	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)//decodujesz z cursora i wrzucasz do modelu
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//########################################################################################
//########################################################################################

func (db *DB) DeleteOneRecord(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deleteCount, err := db.collection.DeleteOne(context.Background(),filter)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("modified count", deleteCount.DeletedCount)
}

//########################################################################################
//########################################################################################

func (db *DB) DeleteAllRecords() {
	deleteCount, err := db.collection.DeleteMany(context.Background(), bson.M{},nil) //select all {{}} to znacczy

	if err != nil {
		log.Fatal(err) 
	}
	log.Println("modified count", deleteCount.DeletedCount) 
}

//########################################################################################
//########################################################################################


func (db *DB) UpdateOneRecord(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID) //dostajemy id do mongodb, ze stringa movieID
	filter := bson.M{"_id": id} //to chyba szuka obiektu po id
	update := bson.M{"$set": bson.M{"watched":true}}//a tu ustawiamy setter, zamieniajacy property "watched" dla danego obiektu

	result, err := db.collection.UpdateOne(context.Background(), filter, update)//i chyba tutaj to wszystko wykonujemy
	if err != nil {
		log.Fatal(err)
	}
	log.Println("modified count", result.ModifiedCount)

}
