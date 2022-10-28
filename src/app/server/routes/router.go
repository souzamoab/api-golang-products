package routes

import (
	"github.com/fabriciojlm/api-golang-users/src/app/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes() *gin.Engine {
	router := gin.Default()

	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowAllUsers)
			users.GET("/health", controllers.Health)
			users.GET("/:id", controllers.GetUserById)
			users.POST("/", controllers.CreateUser)
			users.POST("/createWithArray", controllers.CreateUserArray)
			users.PUT("/:id", controllers.EditUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
	return router
}
