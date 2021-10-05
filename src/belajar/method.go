package main

import "fmt"

//membuat property
type User struct {
	ID        int
	firstName string
	lastName  string
	email     string
	isActive  bool
}

type Group struct {
	name        string
	admin       User
	Users       []User
	isAvailable bool
}

func main() {
	//inisialisasi

	user3 := User{20202212, "Ibnu", "Sahal", "hahawsqswh@mail.com", true}

	user3.display(user3)
	// user4 := User{23243243, "Mas", "Sahl", "haewdwsqswh@mail.com", true}

	// users := []User{user3, user4}
	// group1 := Group{"Gamer", user4, users, true}

	// //output displayGroup
	// displayGroup(group1)

}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("User Name :")
	var i int
	for _, user := range group.Users {
		i += 1
		fmt.Println(i, user.firstName, user.lastName)

	}
}
