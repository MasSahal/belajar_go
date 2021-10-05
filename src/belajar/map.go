package main

import "fmt"

func main() {

	myMap := map[string]string{

		// "key" : "value"
		"name":     "Sahal",
		"fakultas": "Fakultas Teknologi dan Informasi",
		"prodi":    "Teknik Informatika",
	}

	for key, i := range myMap {
		fmt.Println("key :", key, "| value :", i)
	}
}
