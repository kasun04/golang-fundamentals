package main

import "fmt"

func main() {
	marks := []float64{87, 94, 84, 90}

	// sum := 0.0
	// for _, val := range marks {
	// 	sum += val
	// }

	fmt.Println(average(marks))

	lat, long, isFound := gpsCodes("San Jose")
	fmt.Println(lat, long, isFound)

	fmt.Println(sum(4, 5, 3, 1, 2, 5))

	fmt.Println(add(4, 5, true))

	// Clousure
	var i float32 = 1.0
	expoIncr := func(factor float32) float32 {
		i += i * factor
		return i
	}

	fmt.Println(expoIncr(2.1))
	fmt.Println(expoIncr(1.1))
	fmt.Println(expoIncr(3.1))

	// Recursion
	fmt.Println(factorial(4))

	// Defer
	// This function will be executed only after the completion of main function
	defer firstFn()
	secondFn()

	// Panic

	dispatch("FOO")
	dispatch("BAR")

	// Pointers
	n := 10
	makeZero(n)
	println("n = ", n)
	makeZeroPtr(&n)
	println("n = ", n)

	// Pointers with 'new'
	mPtr := new(int)
	*mPtr = 9
	println("m = ", *mPtr)
	makeZeroPtr(mPtr)
	println("m = ", *mPtr)
}

func average(values []float64) float64 {
	total := 0.0
	for _, val := range values {
		total += val
	}
	return total / float64(len(values))
}

//multi param
func add(x, y int, status bool) int {
	fmt.Println(status)
	return x + y
}

// Multi Return
func gpsCodes(location string) (latitude float64, longitude float64, status bool) {

	if location == "San Jose" {
		latitude = 37.494473
		longitude = -122.074585
		status = true
	} else {
		longitude = -1
		latitude = -1
		status = false
	}
	return
}

// Variadic Functions
func sum(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

// Closure : Even Num Generator

// Recursion
func factorial(val int) int {
	if val == 0 {
		return 1
	}
	return val * factorial(val-1)
}

// Defer
func firstFn() {
	fmt.Println("Func : 1")
}

func secondFn() {
	fmt.Println("Func : 2")
}

// Panic
func dispatch(msg string) {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println("Recovered from Panic -> ", str)
		} else {
			fmt.Println("No panic occured!")
		}

	}()
	if msg == "FOO" {
		fmt.Println(msg + " : No Panic")
	} else {
		// Panic indicates a programmer error
		panic("Not received a FOO message. -> Panic")
	}
}

func makeZero(x int) {
	x = 0
}

func makeZeroPtr(xPtr *int) {
	*xPtr = 0
}
