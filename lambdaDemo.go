package main

import "fmt"

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	accumulate := Accumulate(1)

	fmt.Println(accumulate())

	fmt.Println(accumulate())

	fmt.Printf("%p\n", &accumulate)

	accumulate1 := Accumulate(10)

	fmt.Println(accumulate1())

	fmt.Printf("%p\n", &accumulate1)
}
