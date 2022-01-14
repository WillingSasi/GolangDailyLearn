package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.RWMutex

func main() {
	// m = new(sync.RWMutex)

	go read(1)
	go read(2)
	time.Sleep(2 * time.Second)
}

func read(i int) {
	fmt.Println(i, "read start")
	m.RLock()
	fmt.Println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	fmt.Println(i, "read over")
}
