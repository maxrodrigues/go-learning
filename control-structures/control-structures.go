package main

import "fmt"

func main() {
	number := 10
	if number > 15 {
		fmt.Println("Greater then 15")
	} else {
		fmt.Println("Less then 15")
	}

	// iniciando uma variavel e realizando a comparação (if init)
	// e a variavel é somente do escopo do if
	if numberTwo := number; numberTwo > 0 {
		fmt.Println("Number greather then 0")
	}
}
