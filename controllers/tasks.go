package controllers

import (
	"fmt"

	"github.com/heroku/go-getting-started/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTask(c *gin.Context) {
	var task []models.Task
	
	result := db.Find(&task).Order("name ASC")

	if result.Error == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to get all tasks"})
	}
}

func GetOneTask(c *gin.Context) {
	var task models.Task
	
	id := c.Params.ByName("id")
	result := db.First(&task, id)

	if result.Error == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to get one task"})
	}
}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)

	result := db.Create(&task)
	if result.Error == nil {
		c.JSON(201, task);
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to create task"})
	}
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)
	
	id := c.Params.ByName("id")
	db.Model(&models.Task{}).Where("id = ?", id).Updates(
		models.Task{Priority: task.Priority, TaskStatus: task.TaskStatus})
	
	result := db.Save(&task)
	
	if result.Error == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to update task"})
	}
}

func EditTask(c *gin.Context) {
	var task models.Task
	c.Bind(&task)
	
	id := c.Params.ByName("id")
	db.Model(&models.Task{}).Where("id = ?", id).Updates(
		models.Task{
			TaskName   : task.TaskName,
			Details    : task.Details,
			TagId      : task.TagId,
			Deadline   : task.Deadline,
			Priority   : task.Priority,
			TaskStatus : task.TaskStatus,
			CreatedBy  : task.CreatedBy,
			AssignedTo : task.AssignedTo,
		})
	result := db.Save(&task)
	
	if result.Error == nil {
		c.JSON(200, task)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to edit task"})
	}
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Params.ByName("id")

	result := db.Delete(&task, id)
	
	if result.Error == nil {
		c.JSON(200, task)
	} else {
		c.JSON(404, gin.H{"error": "failed to delete task"})
	}
}
