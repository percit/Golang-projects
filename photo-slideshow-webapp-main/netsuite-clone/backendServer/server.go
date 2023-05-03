package main

import (
	"log"
	"backendServer/database"
	"net/http"
	"github.com/gin-gonic/gin"
	"context"
	"golang.org/x/sync/errgroup"
)

type Routes struct {
	Database database.DB
}

func main() {
	db := database.DB{}
	db.InitDB()

	routes := Routes{
		Database: db,
	}
	r := gin.Default()
	r.GET("/ping", routes.ping)
	r.GET("/api/hours", routes.getHours)
	r.POST("/api/hours", routes.setHours)

	errs, _ := errgroup.WithContext(context.Background())

	errs.Go(func() error {
		err := r.Run()
		return err
	})

	err := errs.Wait()
	log.Fatal(err)
}
func (r *Routes) ping(c *gin.Context) {
	log.Println("ping")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (r *Routes) getHours(c *gin.Context) {
	log.Println("getHours")
	todayDate := "24.04.2023"//tu by bylo trzeba zamienic to
	p := r.Database.GetHoursByDate(todayDate)
	c.JSON(http.StatusOK, p)
}

func (r *Routes) setHours(c *gin.Context) {
	log.Println("setHours")
	todayDate := "24.04.2023"//tu by bylo trzeba zamienic to
	todayHours := 8
	r.Database.SetHoursByDate(todayDate, todayHours)
	c.Status(201)
}


//dodaje funkcje, wywolujace funkcje z db.go
//login (ale to moze)
//autoryzacja (tak samo jak z loginem, to juz na samym koncu)
