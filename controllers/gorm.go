package controllers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = initDb()

func initDb() *gorm.DB {

	host     := os.Getenv("HOST")
	user     := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname   := os.Getenv("DBNAME")
	endpoint := os.Getenv("ENDPOINT")
	psqlInfo := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=require&options=endpoint%%3D%s",
		user, password, host, dbname, endpoint,
	)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://task-management-yixiann.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
