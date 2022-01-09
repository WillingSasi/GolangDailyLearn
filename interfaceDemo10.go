package main

import "fmt"

func main() {
	var any interface{}

	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)

	var a int = 1

	var i interface{} = a
	var b int = i.(int)

	fmt.Println(b)

	var x interface{} = 100
	var y interface{} = "hi"

	fmt.Println(x == y)

	// var c interface{} = []int{10}
	// var d interface{} = []int{20}

	// fmt.Println(c == d)
}
