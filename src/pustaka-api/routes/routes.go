package routes

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	//mirip route group
	v1 := router.Group("/v1")

	//api versi 1
	v1.GET("/", handler.RootHandler)
	v1.POST("/books", handler.PostInputBooks)

	//api versi 2
	v2 := router.Group("/v2")
	v2.GET("/", handler.RootHandler2)
	return router

}
