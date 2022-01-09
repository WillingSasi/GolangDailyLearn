package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3)
	printSlice(numbers)

	numbers = append(numbers, 4, 5)
	printSlice(numbers)

	// numbers = append(numbers, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)
	// printSlice(numbers)

	// numbers1 := make([]int, len(numbers), cap(numbers)*2)
	// printSlice(numbers1)

	// copy(numbers1, numbers)
	// printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
