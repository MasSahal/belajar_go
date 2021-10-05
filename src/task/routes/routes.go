// routes/routes.go
package routes

import (
	"task/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	route.GET("/tasks", controllers.FindTasks)
	route.POST("/tasks", controllers.CreateTask)
	route.GET("/tasks/:id", controllers.FindTask)
	route.PUT("/tasks/:id", controllers.UpdateTask)
	route.DELETE("tasks/:id", controllers.DeleteTask)
	return route
}
