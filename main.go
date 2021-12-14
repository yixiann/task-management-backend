package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/heroku/go-getting-started/mappings"
)

func main() {

	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run("localhost:8080")
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// 	_ "github.com/heroku/x/hmetrics/onload"
// )

// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.LoadHTMLGlob("templates/*.tmpl.html")
// 	router.Static("/static", "static")

// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl.html", nil)
// 	})

// 	router.Run(":" + port)

// }
