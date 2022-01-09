package main

import (
	"fmt"
)

const MAX int = 3

func main() {
	// a := 10
	// b := a

	// fmt.Println(&a)
	// fmt.Println(&b)

	var a int = 20
	var ip *int
	var ptr *int

	ip = &a
	fmt.Printf("a变量的地址是: %x\n", &a)

	fmt.Printf("ip变量存储的指针地址: %x\n", ip)

	fmt.Printf("*ip变量的值：%d\n", *ip)

	fmt.Printf("ptr的值为: %x\n", ptr)

	if ptr == nil {
		fmt.Printf("ptr的值为nil: %x\n", ptr)
	}

	b := []int{10, 100, 100}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, b[i])
	}

	var ptr1 [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr1[i] = &b[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr1[i])
	}
}
