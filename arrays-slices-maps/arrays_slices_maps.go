package main

import "fmt"

func main() {

	// ********** Arrays **********

	var x [5]int
	x[4] = 999
	fmt.Println(x)

	var marks [5]float64
	marks[0] = 98
	marks[1] = 87
	marks[2] = 97
	marks[3] = 88
	marks[4] = 81

	var total float64 = 0
	for i := 0; i < len(marks); i++ {
		total += marks[i]
	}
	fmt.Println(total / float64(len(marks)))

	var sum float64 = 0

	// Iterators value is ignored with '_'
	for _, value := range marks {
		sum += value
	}
	fmt.Println(sum / float64(len(marks)))

	marks2 := [5]float64{99,
		88,
		87,
		65,
		90,
	}

	sum = 0
	for _, value := range marks2 {
		sum += value
	}
	fmt.Println(sum / float64(len(marks2)))

	// ********** Slices **********
	mySlice := make([]float64, 5, 10)
	mySlice[0] = 99
	mySlice[1] = 88.3
	fmt.Println(mySlice)

	// Slice with limits and markers
	slice1 := marks2[0:2]
	fmt.Println(slice1)
	slice1 = marks2[2:4]
	fmt.Println(slice1)

	// Slice : Append

	// append creates a new slice by augmenting the exiting content.
	slice2 := append(slice1, 78, 69)
	fmt.Println(slice1, slice2)

	// Slice : Copy
	sliceX := []int{1, 2, 3}
	sliceY := make([]int, 10)

	// dest, src
	copy(sliceY, sliceX)

	fmt.Println(sliceX, sliceY)

	// ******* Maps *********

	// RT Error: maps have to be initialize before they used.
	// var marksMap map[string]float64

	marksMap := make(map[string]float64)
	marksMap["alex"] = 99.8
	marksMap["bob"] = 87.0
	marksMap["katie"] = 83.9
	marksMap["alice"] = 80.9

	fmt.Println(marksMap)

	// Multi-returns
	katiesMarks, status := marksMap["katie"]
	fmt.Println(katiesMarks)
	fmt.Println(status)

}
