package main

import "fmt"

type user struct {
	name    string
	age     uint
	address address
}

type address struct {
	street string
	number int
}

func main() {
	var userOne user
	userOne.name = "Maxuel Barbosa Rodrigues"
	userOne.age = 38
	fmt.Println(userOne)

	address := address{"Beco Aurilio Butter", 46}
	userTwo := user{"Ranis", 35, address}
	fmt.Println(userTwo)

	userThree := user{name: "Jujuba"}
	fmt.Println(userThree)
}
