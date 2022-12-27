package controllers

import (
	"fmt"

	"github.com/heroku/go-getting-started/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTag(c *gin.Context) {
	var tag []models.Tag
	
	result := db.Find(&tag)

	if result.Error == nil {
		c.JSON(200, tag)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to get all tags"})
	}
}

func CreateTag(c *gin.Context) {
	var tag models.Tag
	c.Bind(&tag)

	result := db.Create(&tag) 
	
	if result.Error == nil {
		c.JSON(201, tag);
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to create tag"})
	}
}

func EditTag(c *gin.Context) {
	var tag models.Tag
	c.Bind(&tag)
	
	id := c.Params.ByName("id")
	db.Model(&models.Tag{}).Where("id = ?", id).Updates(
		models.Tag{TagName: tag.TagName, Colour: tag.Colour})
	
	result := db.Save(&tag)
	
	if result.Error == nil {
		c.JSON(200, tag)
	} else {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "failed to edit tag"})
	}
}

func DeleteTag(c *gin.Context) {
	var tag models.Tag

	id := c.Params.ByName("id")
	result := db.Delete(&tag, id)

	if result.Error == nil {
		c.JSON(200, tag)
	} else {
		c.JSON(404, gin.H{"error": "failed to delete tag"})
	}
}
