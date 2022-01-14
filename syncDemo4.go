package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.RWMutex

func main() {
	go write(1)
	go read(2)
	go write(3)
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

func write(i int) {
	fmt.Println(i, "write start")
	m.Lock()
	fmt.Println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	fmt.Println(i, "write over")
}
