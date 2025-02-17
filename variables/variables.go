package main

import "fmt"

func main() {
	// There is the possibility of declaring a variable,
	// including its type.
	var firstVariable string = "First variable"
	fmt.Println(firstVariable)

	// One possibility is type inference, that is, the
	// language determines the type of the variable without
	// explicitly declaring it.
	secondVariable := "Second variable"
	fmt.Println(secondVariable)

	// It is also possible to declare several variables at once.

	// Declaring multiple variables and declaring their types
	var (
		thirdVariable  string = "Third Variable"
		fourthVariable string = "Fourth Variable"
	)

	fmt.Println(thirdVariable, fourthVariable)

	// There is the possibility of inferring types of multiple
	// variables as well.
	fifthVariable, sixthVariable := "Fifth Variable", "Sixth Variable"
	fmt.Println(fifthVariable, sixthVariable)

	// The same principle can be applied to constants
}
