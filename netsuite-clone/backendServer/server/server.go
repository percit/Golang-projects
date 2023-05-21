package server

import (
	"fmt"
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

func (r *Routes) InsertRecord(c *gin.Context) {
	c.Header("Content-Type", "application/x-www-form-urlencoded")
	c.Header("Allow-Control-Allow-Methods", "POST")

	var model model.User
	if err := c.ShouldBindJSON(&model); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	r.Database.InsertUser(model)
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
	
	r.Database.DeleteAllRecords()
	c.JSON(http.StatusOK, gin.H{"message":"All records deleted"})
}

func (r *Routes) UpdateOneRecord(c *gin.Context) {
	fmt.Println("test123")
	c.Header("Content-Type", "application/x-www-form-urlencode")
	c.Header("Allow-Control-Allow-Methods", "PUT")

	id := c.Param("id")
	fmt.Println(id)
	r.Database.UpdateOneRecord(5, id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated one record", "id": id})
}