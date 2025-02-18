package main

import "fmt"

func main() {
	fmt.Println(":-----------------------:")
	fmt.Println("Array")
	// In go the array has fixed positions and types, an
	// array of 3 integer positions cannot receive a
	// string, nor have a fourth position.
	var arrayOne [5]int
	arrayOne[0] = 10
	fmt.Println(arrayOne)

	arrayTwo := [5]string{"One", "Two", "Three", "Four", "Five"}
	fmt.Println(arrayTwo)
	// arrayTwo[5] = "Six" <- argument invalid

	fmt.Println(":-----------------------:")
	fmt.Println("Slices")

	// The only thing that differentiates a slice from an
	// array is that it has "dynamic positions", but as
	// for the type, it remains immutable.
	sliceOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sliceOne)

	// Add new value
	sliceOne = append(sliceOne, 123)
	fmt.Println(sliceOne)

	// The slice is created from an array (not necessarily
	// declared), and it is possible to create slices from
	// an array, just as slices created from arrays can be
	// affected when the array is changed.

	// Creating slice from array
	sliceTwo := arrayTwo[1:4]
	fmt.Println(sliceTwo)

	arrayTwo[3] = "Four - Three"
	fmt.Println(sliceTwo)
	fmt.Println(arrayTwo)
}
