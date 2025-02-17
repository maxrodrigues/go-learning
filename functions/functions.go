package main

import "fmt"

func main() {
	// Calling the sum function
	sumResult := sum(10, 5)
	fmt.Println(sumResult)

	// Another way would be to create a function type variable
	var function = func() {
		fmt.Println("Print text here")
	}
	function()

	// We can pass parameters
	var functionTwo = func(txt string) {
		fmt.Println(txt)
	}
	functionTwo("Text in function declaration")

	// It is also possible to add returns
	var functionThree = func(txt string) int {
		fmt.Println(txt)
		return 15
	}
	functionThree("Another text in function declaration")

	// In order to be used without problems, the result of
	// the function must be assigned to variables according
	// to the return. If a function has two returns, it
	// must have the return assignment in two variables.
	// If it has three, in three, and so on.
	sumMathResult, subsMathResult := mathCalc(5, 8)
	fmt.Println(sumMathResult, subsMathResult)

	// If I don't want one of the returns, I can replace
	// the variable with a _ so it will ignore the return
	// that I don't want to display or use.
	sumMathResultTwo, _ := mathCalc(5, 8)
	fmt.Println(sumMathResultTwo)

	_, subsMathResultTwo := mathCalc(5, 8)
	fmt.Println(subsMathResultTwo)
}

// This would be the basic way to make a function
// and call it.
func sum(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

// Go also supports more than one return in the function,
// this is very common to add error returns
func mathCalc(numberOne, numberTwo int) (int, int) {
	sum := numberOne + numberTwo
	subs := numberOne - numberTwo

	return sum, subs
}
