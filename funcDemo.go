package main

import "fmt"

func main() {
	// f := func(data int) {
	// 	fmt.Println("hello", data)
	// }

	// f(100)
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})
}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
