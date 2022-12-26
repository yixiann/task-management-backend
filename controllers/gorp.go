package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {

	host     := os.Getenv("HOST")
	port     := os.Getenv("PORT")
	user     := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname   := os.Getenv("DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+ "password=%s dbname=%s sslmode=enable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}

	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "https://task-management-yixiann.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
