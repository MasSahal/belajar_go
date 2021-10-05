package book

import "time"

//struct itu mirip model
type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Discount    int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
