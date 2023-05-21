package main

import (
	// "log"
	// "context"
	"flag"

	"backendServer/database"
	"backendServer/database/model"
	// "backendServer/server"

	// "github.com/gin-gonic/gin"
	// "golang.org/x/sync/errgroup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MongoKey string

func init() {
	flag.StringVar(&MongoKey, "m", "", "MongoDB Key")
	flag.Parse()
}

func main() {
	db := database.DB{}
	db.InitDB(MongoKey)

	user1 := model.User{
		ID: primitive.NewObjectID(),
		WorkerName: "John Kowalski",
		Date: "22.05.2022",
		Hours: 8}	
	user2 := model.User{
		ID: primitive.NewObjectID(),
		WorkerName: "Piotr Lukasiewicz",
		Date: "23.05.2022",
		Hours: 6}	
	user3 := model.User{
		ID: primitive.NewObjectID(),
		WorkerName: "Barack Obama",
		Date: "24.05.2022",
		Hours: 0}	

	db.InsertUser(user1)
	db.InsertUser(user2)
	db.InsertUser(user3)

	db.UpdateOneRecord(5, user2.ID) //idk it this works

	db.FindUser("John Kowalski") //idk it this works

	db.DeleteOneRecord(user1.ID)

	db.DeleteAllRecords()

	db.CloseConnection()








	// routes := server.Routes{
	// 	Database: db,
	// }
	// r := gin.Default()
	// r.GET("/ping", routes.Ping)
	// r.GET("/api/GetHours", routes.GetHours)
	// r.POST("/api/SetHours", routes.SetHours)

	// r.GET("/api/GetAllRecords", routes.GetAllRecords)
	// r.POST("/api/InsertRecord/{id}", routes.InsertRecord)
	// r.PUT("/api/UpdateOneRecord/{id}", routes.UpdateOneRecord)
	// r.DELETE("/api/DeleteOneRecord{id}", routes.DeleteOneRecord)
	// r.DELETE("/api/DeleteAllRecords", routes.DeleteAllRecords)

	// errs, _ := errgroup.WithContext(context.Background())

	// errs.Go(func() error {
	// 	err := r.Run()
	// 	return err
	// })

	// err := errs.Wait()
	// log.Fatal(err)
}


// TODO
// - stworz funkcje, ktore dodaja do bazy danych zabite wartosci przy POST itd, nie probuj wyciagnac ich z api
// - daj jakis przyklad, ze wywolasz 3 razy taka sama funkcje najpierw z roznymi wartosciami, przed wywolaniem GET