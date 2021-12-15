package controllers

import (
	"fmt"

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
	// INSERT INTO `tags` (`id`, `tag_name`, `colour`) VALUES ('3', 'Sleep', 'cyan');

	if tag.TagName != "" || tag.Colour != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO tags (tag_name, colour) VALUES (?, ?)`,
			tag.TagName, tag.Colour); insert != nil {
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

func EditTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var tag models.Tag
	c.Bind(&tag)

	// UPDATE `tags` SET `tag_name` = 'Eating', `colour` = 'red' WHERE `tags`.`id` = 104

	if tag.TagName != "" || tag.Colour != "" {
		_, err := dbmap.Exec("UPDATE tags SET tag_name=?, colour=? WHERE id=?",
			tag.TagName, tag.Colour, id)
		if err == nil {
			content := &models.Tag{
				Id:      tag.Id,
				TagName: tag.TagName,
				Colour:  tag.Colour,
			}
			c.JSON(201, content)
		} else {
			checkErr(err, "Edit failed")
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

	// update,_ := dbmap.Update(&tag, "UPDATE tags SET tag_name=?, colour=? WHERE id=?",
	// 	tag.TagName, tag.Colour, id)
	// if err == nil {
	// 	c.JSON(200, "Success")
	// } else {
	// 	c.JSON(404, gin.H{"error": "tag not found"})
	// }
}

func DeleteTag(c *gin.Context) {
	id := c.Params.ByName("id")
	var tag models.Tag
	_, err := dbmap.Select(&tag, "DELETE FROM tags WHERE id=?", id)
	if err == nil {
		c.JSON(200, "Success")
	} else {
		c.JSON(404, gin.H{"error": "tag not found"})
	}
}
