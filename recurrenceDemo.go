package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}

	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	i := 15

	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	for j := 0; j < 20; j++ {
		fmt.Printf("%d 的斐波那契数是 %d\n", j, fibonacci(j))
	}
}
