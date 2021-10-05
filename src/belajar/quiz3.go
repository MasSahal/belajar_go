package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := calculate(20, 10, "4")

	if err != nil {
		fmt.Println("Terjadi kesalahan fatal")
		fmt.Println(err.Error())
	}

	if result != 0 {
		fmt.Println(result)
	}
}

func calculate(number, numberTwo int, operation string) (int, error) {

	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo

	case "-":
		result = number - numberTwo

	case "*":
		result = number * numberTwo

	case "/":
		result = number / numberTwo

	default:
		errorResult = errors.New("Unknown Operation!")
	}

	return result, errorResult
}

func sum(numbers []int) int {
	var total int

	for _, number := range numbers {
		total = total + number
	}
	return total
}

// mencari rata rata
// scores := [8]int{100, 80, 45, 93, 98, 87, 77, 70}
// scores := [8]int{100, 100, 100, 100, 50, 50, 50, 73}

// var total int

// for _, score := range scores {

// 	total = total + score
// }

// lenght := len(scores)
// //kasih float64 biar bisa ada pecahan 0.0000
// average := float64(total) / float64(lenght)

// fmt.Println(average)

// benar := false
// salah := true

// if benar == false {
// 	benar = false
// 	benarJuga := benar

// 	fmt.Println(benarJuga)

// } else if salah == true {
// 	salah = false
// 	salahJuga := salah

// 	fmt.Println(salahJuga)
// }
