package main

import (
	"errors"
	"fmt"
)

func main() {
	// The declaration of int, can be int8, int16, int32
	// and int64. This declaration is about the number of
	// bits it can occupy. If the simple declaration is
	// made, the type will be according to the architecture
	// of your computer.
	var numberOne int64 = -100000000
	fmt.Println(numberOne)

	// uint works the same way, however, it is used for
	// unsigned numbers
	var numberTwo uint32 = 10000
	fmt.Println(numberTwo)

	// alias
	// rune is the same thing as int32, but it is more
	// commonly used when referring to a number that
	// represents a character.
	var numberThree rune = 123444
	fmt.Println(numberThree)

	// bity is the same thing as uint8, since it
	// represents a bit
	var numberFour byte = 123
	fmt.Println(numberFour)

	// Real numbers can be used with float32 or float64
	// declaration, in type inference it will be used
	// according to your computer architecture
	var numberFive float32 = 123.45
	fmt.Println(numberFive)

	var numberSix float64 = 12300000000.45
	fmt.Println(numberSix)

	// For strings always use double quotes, if single
	// quotes are used it will be considered as a char and
	// converted to a number
	var str string = "text here"
	fmt.Println(str)

	str2 := "Another text"
	fmt.Println(str2)

	var boolanOne bool
	fmt.Println(boolanOne)

	// Golang has a specific type, called error, it does
	// not support a string or integer, being its own type
	// and having a native package for its operation.
	var errorOne error = errors.New("Internal Server Error")
	fmt.Println(errorOne)
}
