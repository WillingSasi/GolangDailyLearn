package main

import "fmt"

func main() {
	fmt.Println("defer begin")

	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")
}

var {
	valueByKey = make(map[string]int)
	valueByKeyGuard sync.Mutex
}

func readValue(key string) int {
	valueByKeyGuard.Lock()

	v := valueByKey[key]

	valueByKeyGuard.Unlock()

	return v
}

func readValueDefer(key string) int {
	valueByKeyGuard.Lock()

	 defer valueByKeyGuard.Unlock()

	return valueByKey[key]
}
