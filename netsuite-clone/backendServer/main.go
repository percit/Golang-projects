package main

import (
	"log"
	"context"
	"flag"
	"fmt"

	"backendServer/database"
	"backendServer/database/model"

	"backendServer/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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

	// db.InsertUser(user1)
	// db.InsertUser(user2)
	// db.InsertUser(user3)

	// db.UpdateOneRecord(5, user2.ID)

	// fmt.Println(db.FindUser("John Kowalski"))

	// db.DeleteOneRecord("646a5c795ef8c8ac407e72c8")

	// db.DeleteAllRecords()

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)










	routes := server.Routes{
		Database: db,
	}
	r := gin.Default()
	r.GET("/ping", routes.Ping) //works
	r.GET("/api/GetAllRecords", routes.GetAllRecords)//works
	r.POST("/api/InsertRecord", routes.InsertRecord) //works (this needs to be in Form, not Json Payload)
	r.PUT("/api/UpdateOneRecord/:id", routes.UpdateOneRecord) //works
	r.DELETE("/api/DeleteOneRecord/:id", routes.DeleteOneRecord)//works
	r.DELETE("/api/DeleteAllRecords", routes.DeleteAllRecords) //works

	errs, _ := errgroup.WithContext(context.Background())

	errs.Go(func() error {
		err := r.Run()
		return err
	})

	err := errs.Wait()
	log.Fatal(err)

	db.CloseConnection()
}