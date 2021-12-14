package controllers

import (
	"database/sql"
	// "fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/heroku/go-getting-started/models"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:Welcome1!@tcp(127.0.0.1:3306)/task_management")

	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()

	checkErr(err, "sql.Open failed")
	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap := &gorp.DbMap{Db: db}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap

	// results, err := db.Query("SELECT * FROM users")

	// insert, err := db.Query("INSERT INTO users VALUES('2','test', 'test','test','test')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var user models.User

	// 	err = results.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname, &user.Lastname)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Println(user.Username)

	// }

	// defer insert.Close()
	// fmt.Println("SUCCESS MAYBE")
	// fmt.Println("SUCCESS YAY")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
