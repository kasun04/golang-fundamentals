package main

import "fmt"

func tryConditions() {
	i := 7

	// if, if-else, else
	if i%2 == 0 {
		fmt.Println("Divisible by 2")
	} else if i%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Divisible by other number")
	}

	// Swtich
	j := 2

	switch j {
	case 0:
		fmt.Println("Case 0")
		j = -1
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		j = j + 10
	}

	fmt.Println(j)
}
