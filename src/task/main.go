package main

import (
	// book adalah directory root project go yang kita buat
	"task/models" // memanggil package models pada directory models
	"task/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
