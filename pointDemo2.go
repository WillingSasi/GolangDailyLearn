package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)

	swap2(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func swap2(x *int, y *int) {
	*x, *y = *y, *x
}
