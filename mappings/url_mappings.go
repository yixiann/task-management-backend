package mappings

import (
	"github.com/heroku/go-getting-started/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Router.Use(controllers.Cors())

	user := Router.Group("/user")
	{
		user.GET("/users/:id", controllers.GetUserDetail)
		user.GET("/users/", controllers.GetUser)
		user.POST("/login/", controllers.Login)
		user.PUT("/users/:id", controllers.UpdateUser)
		user.POST("/users", controllers.PostUser)
	}

	task := Router.Group("/task")
	{
		task.GET("/fetch", controllers.GetAllTask)
		task.GET("/details/:id", controllers.GetOneTask)
		// task.POST("/create", controllers.CreateTask)
		// task.PUT("/update/:id", controllers.UpdateTask)
		// task.PUT("/edit/:id", controllers.EditTask)
		// task.DELETE("/delete/:id", controllers.DeleteTask)
	}

	tag := Router.Group("/tag")
	{
		tag.GET("/fetch", controllers.GetAllTag)
		tag.POST("/create", controllers.CreateTag)
		// tag.GET("/edit/:id", controllers.EditTag)
		// tag.GET("/delete/:id", controllers.DeleteTag)
	}
}
