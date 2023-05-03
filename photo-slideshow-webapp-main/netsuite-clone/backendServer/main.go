package main

import (
	"log"
	"context"

	"backendServer/database"
	"backendServer/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	db := database.DB{}
	db.InitDB()

	routes := server.Routes{
		Database: db,
	}
	r := gin.Default()
	r.GET("/ping", routes.Ping)
	r.GET("/api/hours", routes.GetHours)
	r.POST("/api/hours", routes.SetHours)

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
