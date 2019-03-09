package main

import "fmt"

func main() {

	fmt.Println("1 + 1 = ", 1+1)

	// Lenth of the string
	fmt.Println("Len : ", len("Hello, World"))

	// Index [1] of the string, returns the byte
	fmt.Println("Hello, World"[1])

	// Concat
	fmt.Println("Hello, " + "World")

	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
