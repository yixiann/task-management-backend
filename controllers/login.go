package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/heroku/go-getting-started/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTask(c *gin.Context) {
	var task []models.Task
	_, err := dbmap.Select(&task, "SELECT * FROM tasks")

	if err == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "tasks not found"})
	}
}

func GetUser(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "SELECT * FROM users")

	fmt.Println(user)
	fmt.Println("WHY")

	if err == nil {
		c.JSON(200, user)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func GetUserDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)

	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)

		content := &models.User{
			Id:       user_id,
			Username: user.Username,
			Email:    user.Email,
			TaskId:   user.TaskId,
			// Password:  user.Password,
			// Firstname: json.Firstname,
			// Lastname:  json.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func Login(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)

	if err == nil {
		user_id := user.Id

		content := &models.User{
			Id:       user_id,
			Username: user.Username,
			Email:    user.Email,
			TaskId:   user.TaskId,
			// Password:  user.Password,
			// Firstname: json.Firstname,
			// Lastname:  json.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func PostUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)

	log.Println(user)

	if user.Username != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Username, user.Username, user.Username); insert != nil {
			user_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:       user_id,
					Username: user.Username,
					Email:    user.Email,
					TaskId:   user.TaskId,
					// Password:  user.Password,
					// Firstname: json.Firstname,
					// Lastname:  json.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)

	if err == nil {
		var json models.User
		c.Bind(&json)

		user_id, _ := strconv.ParseInt(id, 0, 64)

		user := models.User{
			Id:       user_id,
			Username: user.Username,
			Email:    user.Email,
			TaskId:   user.TaskId,
			// Password:  user.Password,
			// Firstname: json.Firstname,
			// Lastname:  json.Lastname,
		}

		if user.Username != "" {
			_, err = dbmap.Update(&user)

			if err == nil {
				c.JSON(200, user)
			} else {
				checkErr(err, "Updated failed")
			}

		} else {
			c.JSON(400, gin.H{"error": "fields are empty"})
		}

	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
