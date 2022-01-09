package main

import "fmt"

var a int = 20

func main() {
	a := 10
	b := 20
	c := 0

	fmt.Printf("main()函数中 a = %d\n", a)

	c = sum(a, b)
	fmt.Printf("main()函数中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum()函数中 a = %d\n", a)
	fmt.Printf("sum()函数中 b = %d\n", b)

	return a + b
}
