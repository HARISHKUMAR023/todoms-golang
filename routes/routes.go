package routes

import (
	"todogorest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Todo routes
	todoRoutes := r.Group("/api/v1/todos")
	{
		todoRoutes.GET("/", controllers.GetTodos)
		todoRoutes.GET("/:id", controllers.GetTodoByID)
		todoRoutes.POST("/", controllers.CreateTodo)
		todoRoutes.PUT("/:id", controllers.UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.DeleteTodo)
	}

	return r
}
