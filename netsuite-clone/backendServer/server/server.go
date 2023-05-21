package server

import (
	"log"
	"net/http"

	"backendServer/database"
	"backendServer/database/model"

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

func (r *Routes) InsertRecord(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Allow-Control-Allow-Methods", "POST")

	var model model.Hours
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	r.Database.InsertRecordHelper(model)
	c.JSON(http.StatusOK, gin.H{"message": "Record created", "model": model})
}

func (r *Routes) GetAllRecords(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")

	allModels := r.Database.GetAllRecords()
	c.JSON(http.StatusOK, gin.H{"data": allModels})
}

func (r *Routes) DeleteOneRecord(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Allow-Control-Allow-Methods", "DELETE")

	id := c.Param("id")
	r.Database.DeleteOneRecord(id)
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted", "id": id})
}

func (r *Routes) DeleteAllRecords(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencode")
	c.Header("Allow-Control-Allow-Methods", "DELETE")
	// c.Header("Content-Type", "application/x-www-form-urlencoded")
	// c.Header("Allow-Control-Allow-Methods", "DELETE")
	
	r.Database.DeleteAllRecords()
	c.JSON(http.StatusOK, gin.H{"message":"All records deleted"})
}

func (r *Routes) UpdateOneRecord(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencode")
	c.Header("Allow-Control-Allow-Methods", "PUT")

	id := c.Param("id")
	r.Database.UpdateOneRecord(id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated one record", "id": id})
}