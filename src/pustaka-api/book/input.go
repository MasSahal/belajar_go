package book

import "encoding/json"

// metode post
type InputBooks struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}
