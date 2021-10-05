package main

import (
	"mahasiswa/config"
	"mahasiswa/models"
	"mahasiswa/routes"
)

func main() {

	db := config.SetupDB()
	db.AutoMigrate(&models.Mahasiswa{})

	router := routes.SetupRoutes(db)
	router.Run(":8000")
}
