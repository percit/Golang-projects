package main

import (
	"log"
	"context"
	"flag"

	"backendServer/database"
	"backendServer/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var MongoKey string

func init() {
	flag.StringVar(&MongoKey, "m", "", "MongoDB Key")
	flag.Parse()
}

func main() {
	db := database.DB{}
	db.InitDB(MongoKey)

	routes := server.Routes{
		Database: db,
	}
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	r.GET("/api/GetHours", routes.GetHours)
	r.POST("/api/SetHours", routes.SetHours)

	r.GET("/api/GetAllRecords", routes.GetAllRecords)
	r.POST("/api/InsertRecord/{id}", routes.InsertRecord)
	r.PUT("/api/UpdateOneRecord/{id}", routes.UpdateOneRecord)
	r.DELETE("/api/DeleteOneRecord{id}", routes.DeleteOneRecord)
	r.DELETE("/api/DeleteAllRecords", routes.DeleteAllRecords)

	errs, _ := errgroup.WithContext(context.Background())

	errs.Go(func() error {
		err := r.Run()
		return err
	})

	err := errs.Wait()
	log.Fatal(err)
}


//login (ale to moze)
//autoryzacja (tak samo jak z loginem, to juz na samym koncu)
