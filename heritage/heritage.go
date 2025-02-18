package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint
	height   uint
}

type student struct {
	person
	course  string
	college string
}

func main() {
	firstPerson := person{"Max", "Rodrigues", 38, 177}
	fmt.Println(firstPerson)

	FirstStudent := student{firstPerson, "TADS", "Sale"}
	fmt.Println(FirstStudent)
	fmt.Println(FirstStudent.name)
}
