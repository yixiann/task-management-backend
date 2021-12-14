package controllers

import (
	"fmt"
	"log"

	"github.com/heroku/go-getting-started/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllTag(c *gin.Context) {
	var tag []models.Tag
	_, err := dbmap.Select(&tag, "SELECT * FROM tags")

	if err == nil {
		c.JSON(200, tag)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "tags not found"})
	}
}

func CreateTag(c *gin.Context) {
	var tag models.Tag
	c.Bind(&tag)
	log.Println(tag)
	if tag.TagName != "" || tag.Colour != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO tags (Id, TagName, TagColour) VALUES (?, ?, ?)`,
			tag.Id, tag.TagName, tag.Colour); insert != nil {
			tag_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.Tag{
					Id:      tag_id,
					TagName: tag.TagName,
					Colour:  tag.Colour,
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
