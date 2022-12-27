package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = initDb()

func initDb() *gorm.DB {

	// host     := os.Getenv("HOST")
	// user     := os.Getenv("USER")
	// password := os.Getenv("PASSWORD")
	// dbname   := os.Getenv("DBNAME")
	// port 		 := os.Getenv("PORT")
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	DEFAULT_PORT := 5432
  
  host     := "database-yixian.cwmt6ozdpjrl.ap-southeast-1.rds.amazonaws.com"
  user     := "postgresYiXian"
  password := "YiXianPostgres"
  dbname   := "database-yixian"

  port, err := strconv.Atoi("5432")
  if err != nil {
    port = DEFAULT_PORT
  }

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
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
