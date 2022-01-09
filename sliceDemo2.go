package main

import "fmt"

func main() {
	var slice [][]int
	slice = [][]int{{10}, {100, 200}}

	fmt.Println(slice)

	slice[0] = append(slice[0], 20)

	fmt.Println(slice)
}
