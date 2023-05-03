package server

import (
	"log"
	"backendServer/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	Database database.DB
}

func (r *Routes) Ping(c *gin.Context) {
	log.Println("ping")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (r *Routes) GetHours(c *gin.Context) {
	log.Println("getHours")
	todayDate := "24.04.2023"//tu by bylo trzeba zamienic to
	p := r.Database.GetHoursByDate(todayDate)
	c.JSON(http.StatusOK, p)
}

func (r *Routes) SetHours(c *gin.Context) {
	log.Println("setHours")
	todayDate := "24.04.2023"//tu by bylo trzeba zamienic to
	todayHours := 8
	r.Database.SetHoursByDate(todayDate, todayHours)
	c.Status(201)
}
