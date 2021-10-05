package routes

import (
	"mahasiswa/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//root index
	route.GET("/mahasiswa", controllers.GetAllMahasiswa)
	route.POST("/mahasiswa", controllers.AddMahasiswa)
	route.GET("/mahasiswa/:id", controllers.GetMahasiswa)
	route.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswa)
	route.PUT("/mahasiswa/:id", controllers.UpdateMahasiswa)
	// route.POST("/data", controllers.DeleteMahasiswa)
	return route
}
