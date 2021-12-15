package controllers

import (
	"fmt"
	"log"

	// "strconv"

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

func GetOneTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	err := dbmap.SelectOne(&task, "SELECT * FROM tasks WHERE id=? LIMIT 1", id)

	if err == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "tasks not found"})
	}
}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)
	log.Println(task)
	if task.TaskName != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO tasks (
			user_id,
			task_name,
			details,
			tag_id,
			deadline,
			priority,
			task_status,
			created_by,
			assigned_to) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			task.UserId,
			task.TaskName,
			task.Details,
			task.TagId,
			task.Deadline,
			task.Priority,
			task.TaskStatus,
			task.CreatedBy,
			task.AssignedTo); insert != nil {
			task_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.Task{
					Id:         task_id,
					UserId:     task.UserId,
					TaskName:   task.TaskName,
					Details:    task.Details,
					TagId:      task.TagId,
					Deadline:   task.Deadline,
					Priority:   task.Priority,
					TaskStatus: task.TaskStatus,
					CreatedBy:  task.CreatedBy,
					AssignedTo: task.AssignedTo,
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

func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task models.Task
	_, err := dbmap.Select(&task, "DELETE FROM tasks WHERE id=?", id)
	if err == nil {
		c.JSON(200, "Success")
	} else {
		c.JSON(404, gin.H{"error": "task not found"})
	}
}

// package controllers

// import (
// 	"log"
// 	"strconv"

// 	"goapi/models"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func GetUser(c *gin.Context) {
// 	var user []models.User
// 	_, err := dbmap.Select(&user, "select * from user")

// 	if err == nil {
// 		c.JSON(200, user)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}

// }

// func GetUserDetail(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)

// 	if err == nil {
// 		user_id, _ := strconv.ParseInt(id, 0, 64)

// 		content := &models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: user.Firstname,
// 			Lastname:  user.Lastname,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }

// func Login(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)
// 	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)

// 	if err == nil {
// 		user_id := user.Id

// 		content := &models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: user.Firstname,
// 			Lastname:  user.Lastname,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}

// }

// func PostUser(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)

// 	log.Println(user)

// 	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {

// 		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Password, user.Firstname, user.Lastname); insert != nil {
// 			user_id, err := insert.LastInsertId()
// 			if err == nil {
// 				content := &models.User{
// 					Id:        user_id,
// 					Username:  user.Username,
// 					Password:  user.Password,
// 					Firstname: user.Firstname,
// 					Lastname:  user.Lastname,
// 				}
// 				c.JSON(201, content)
// 			} else {
// 				checkErr(err, "Insert failed")
// 			}
// 		}

// 	} else {
// 		c.JSON(400, gin.H{"error": "Fields are empty"})
// 	}

// }

// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)

// 	if err == nil {
// 		var json models.User
// 		c.Bind(&json)

// 		user_id, _ := strconv.ParseInt(id, 0, 64)

// 		user := models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: json.Firstname,
// 			Lastname:  json.Lastname,
// 		}

// 		if user.Firstname != "" && user.Lastname != "" {
// 			_, err = dbmap.Update(&user)

// 			if err == nil {
// 				c.JSON(200, user)
// 			} else {
// 				checkErr(err, "Updated failed")
// 			}

// 		} else {
// 			c.JSON(400, gin.H{"error": "fields are empty"})
// 		}

// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
