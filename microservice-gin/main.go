//this wil e article menager, it will handle html, json and xml
package main

import (
	"github.com/gin-gonic/gin"
)
var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*")

   // Handle Index
   router.GET("/", ShowIndexPage) //this is showing all articles on localhost
   router.GET("/article/view/:article_id", getArticle) //this will show only one article
   //to drugie wymaga http://localhost:8080/article/view/1

  // Start serving the application
  router.Run()

}

//zeby to odpalic to musisz dac "go build -o app" i "./app", zwykle go run main.go nie dziala, chyba dlatego, ze nie jest podzielone na foldery

//OLD VERSION

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*") //loading templates html files, you only have to do this once
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML( //call the html method of the Context to render a template
// 			http.StatusOK,
// 			"index.html", //use index.html template
// 			gin.H{
// 				"title": "Home page", //pass data that the page uses
// 			},
// 		)
// 	})
// 	router.Run()
// }
