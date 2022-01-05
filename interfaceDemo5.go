package main

import "fmt"

// func main() {
// 	var x interface{}
// 	x = 10
// 	value, ok := x.(int)

// 	fmt.Print(value, ",", ok)
// }

func main() {
	var a int
	a = 10
	getType(a)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type if a is int")
	case string:
		fmt.Println("the type if a is string")
	case float64:
		fmt.Println("the type if a is float64")
	default:
		fmt.Println("unknow type")
	}
}
