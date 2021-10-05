package main

import "fmt"

func main() {

	// BISA JUGA
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "PS4")
	gamingConsoles = append(gamingConsoles, "PS3")
	gamingConsoles = append(gamingConsoles, "PS2")
	gamingConsoles = append(gamingConsoles, "XBOX360")
	gamingConsoles = append(gamingConsoles, "NINTENDO SWITCH")

	//perulangan mirip foreach
	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

	// var contoh = [...]string{"Nama", "Saya", "Sahal"}

	// fmt.Println(contoh)

}
