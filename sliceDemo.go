package main

import "fmt"

func main() {
	// var numbers = make([]int, 3, 5)
	// printSlice(numbers)

	// var numbers1 []int
	// printSlice(numbers1)

	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers2)
	printSlice(numbers2[1:4])
	printSlice(numbers2[:4])
	printSlice(numbers2[4:])
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
