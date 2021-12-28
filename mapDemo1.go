package main

import (
	"fmt"
	"sync"
)

func main() {

	// m := make(map[int]int)

	// go func() {
	// 	for {
	// 		m[1] = 1
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		_ = m[1]
	// 	}
	// }()

	// for {

	// }

	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	scene.Range(func(key interface{}, value interface{}) bool {
		fmt.Println("iterate: ", key, value)
		return true
	})
}
