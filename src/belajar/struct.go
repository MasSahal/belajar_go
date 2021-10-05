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
	user := User{}

	//mengisi property
	user.ID = 20210001
	user.firstName = "Mas"
	user.lastName = "Sahal"
	user.email = "sahal@localhost.com"
	user.isActive = true

	//bisa juga pake ini
	user2 := User{
		ID:        20202222,
		firstName: "Abu",
		lastName:  "Sahal",
		email:     "hahah@mail.com",
		isActive:  true,
	}

	//bisa juga langsung
	user3 := User{20202212, "Ibnu", "Sahal", "hahawsqswh@mail.com", true}
	user4 := User{23243243, "Mas", "Sahl", "haewdwsqswh@mail.com", true}

	//output displayUser()
	// fmt.Println(displayUser(user))
	// fmt.Println(displayUser(user2))
	// fmt.Println(displayUser(user3))
	// fmt.Println(displayUser(user4))

	users := []User{user, user2, user3, user4}
	group1 := Group{"Gamer", user, users, true}

	//output displayGroup
	displayGroup(group1)

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

//struct sebagai paramter
func displayUser(user User) string {

	//sudah pakai User dari struct, maka bisa di akses di sini
	return fmt.Sprintf("Name : %s %s, Email : %s", user.firstName, user.lastName, user.email)

}
