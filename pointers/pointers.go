package main

import "fmt"

func main() {
	// These variables are copies, and what happens to one
	// does not interfere with the other, if I increase
	// or decrease the value of variableOne, the value of
	// variableTwo does not change.
	var variableOne int = 10
	var variableTwo int = variableOne

	fmt.Println(variableOne, variableTwo)

	variableOne++
	fmt.Println(variableOne, variableTwo)

	// When I assign a variable to a pointer, it stores
	// not the value but the reference in memory.
	var variableThree int
	// You must use * to make the declaration in type
	var pointer *int

	variableThree = 100
	// And the variable to be assigned must have the & in front
	pointer = &variableThree

	fmt.Println(variableThree, pointer)
	fmt.Println(variableThree, *pointer) // derefer
}
