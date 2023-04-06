package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
)

//this file is for validation and data fetching


// Overall, this function determines the format of the response based on the Accept header 
// in the request, and responds with JSON, XML, or HTML accordingly.
func render(c *gin.Context, data gin.H, templateName string) {
  switch c.Request.Header.Get("Accept") {//'Accept' header in the HTTP request that was made to the server
    case "application/json":
      c.JSON(http.StatusOK, data["payload"])// Respond with JSON
    case "application/xml":
      c.XML(http.StatusOK, data["payload"])
    default:
      c.HTML(http.StatusOK, templateName, data)
  }
}
// In practice, gin.H is used as a shorthand way of defining a map of variables that will
//  be passed to an HTML template or used to generate a JSON or XML response. For example, 
//  you might define a gin.H variable like this:
// data := gin.H{
//     "title": "My Website",
//     "body": "Welcome to my website!",
// }



func showIndexPage(c *gin.Context) {
  articles := getAllArticles()

  // Call the render function with the name of the template to render
  render(c, gin.H{
    "title":   "Home Page",
    "payload": articles}, "index.html")

}

func getArticle(c *gin.Context) {
  // Check if the article ID is valid
  if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
    // Check if the article exists
    if article, err := getArticleByID(articleID); err == nil {
      // Call the HTML method of the Context to render a template
      c.HTML(
        // Set the HTTP status to 200 (OK)
        http.StatusOK,
        // Use the index.html template
        "article.html",
        // Pass the data that the page uses
        gin.H{
          "title":   article.Title,
          "payload": article,
        },
      )

    } else {
      // If the article is not found, abort with an error
      c.AbortWithError(http.StatusNotFound, err)
    }

  } else {
    // If an invalid article ID is specified in the URL, abort with an error
    c.AbortWithStatus(http.StatusNotFound)
  }
}