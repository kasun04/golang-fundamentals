package main

import "fmt"

//var globalMsg string = "Hello, Go!"
var globalMsg = "Hello, Go!"

func main() {

	// Variables are created with 'Var' keyword
	// followed by the variable name and type

	var helloMsg string = "Hello, World"
	var greetingMsg string
	greetingMsg = helloMsg + "!"
	fmt.Println(greetingMsg)

	// Shortedn form. Go infers the type of the variable
	msg := "Hello, World"
	fmt.Println(msg)

	// Scoping
	fmt.Println(globalMsg)

	//Constants
	const x int = 7007
	// x = 10 //-> Compiler error
	fmt.Println(x)

	// Multiple variable declaration
	// with var or const keywords
	var (
		a = 10
		b = 4
		c = false
		d = 5.4
		e = "Hello"
	)

	fmt.Println(a, b, c, d, e)

	y := e
	fmt.Println(y + " World")

	z := d
	z = z + 100.1111
	fmt.Println(z)
}
