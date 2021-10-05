package main

import (
	"fmt"
	"pustaka-api/book"
	"pustaka-api/connections"
	"pustaka-api/routes"
)

func main() {

	db := connections.SetupDB()

	//migrate struct ke tabel database
	// db.AutoMigrate(&book.Book{})

	// fmt.Println("Database connection succesed!")

	// book := book.Book{}
	// book.Title = "Pemrograman Web 2"
	// book.Price = 6002300
	// book.Discount = 110
	// book.Rating = 110
	// book.Description = "All in one only one book"

	// err := db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error creating record!")
	// 	fmt.Println("=====================")
	// }

	// var book book.Book
	// err := db.First(&book, 4).Error
	// if err != nil {
	// 	fmt.Println("=====================")
	// 	fmt.Println("Error find record!")
	// 	fmt.Println("=====================")
	// }

	//get all data

	var books []book.Book
	err := db.Debug().Find(&books, 4).Error
	if err != nil {
		fmt.Println("=====================")
		fmt.Println("Error find record!")
		fmt.Println("=====================")
	}

	// for_, b := range books{
	fmt.Println("Title :", books.Title)
	fmt.Println("Price :", books.Price)
	// }

	r := routes.SetupRoutes(db)
	r.Run()
}
