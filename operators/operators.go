package main

import "fmt"

func main() {
	// Arithmetic start
	sum := 1 + 2
	subs := 2 - 1
	division := 10 / 2
	multiplication := 10 * 2
	rest := 10 % 2

	fmt.Println(sum, subs, division, multiplication, rest)

	// When the types of variables are different, as in the
	//  example below, it is not possible to perform
	// arithmetic operations.

	// var numberOne int16 = 10
	// var numberTwo int32 = 15
	// result := numberOne + int16(numberTwo)

	// Arithmetic end

	// Assignment start

	var variableOne string = "Text"
	variableTwo := "Another Text"
	fmt.Println(variableOne, variableTwo)

	// Assignment start

	// relational start
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// relational end

	// logical start
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
	// logical end

	// unary start
	number := 1
	number++
	number += 9
	number--
	number -= 2
	number *= 3
	fmt.Println(number)
	// unary end

	// ternary
	// Ternary operation does not exist in go
}
