package controllers

import (
	"fmt"

	"github.com/heroku/go-getting-started/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllOLD(c *gin.Context) {
	var tag []models.Tag
	_, err := dbmap.Select(&tag, "SELECT * FROM tags")

	if err == nil {
		c.JSON(200, tag)
	} else {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": err})
	}
}
