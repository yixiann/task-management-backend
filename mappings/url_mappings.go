package mappings

import (
	"github.com/heroku/go-getting-started/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())
	// v1 of the API
	task := Router.Group("/task")
	{
		task.GET("/fetch", controllers.GetAllTask)

		task.GET("/users/:id", controllers.GetUserDetail)
		task.GET("/users/", controllers.GetUser)
		task.POST("/login/", controllers.Login)
		task.PUT("/users/:id", controllers.UpdateUser)
		task.POST("/users", controllers.PostUser)
	}
}
