package main

import (
	"fmt"
	"sync"
)

var (
	count int

	// countGuard sync.Mutex
	countGuard sync.RWMutex
)

func GetCount() int {
	// countGuard.Lock()
	countGuard.RLock()

	// defer countGuard.Unlock()
	defer countGuard.RUnlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()

	count = c

	countGuard.Unlock()
}

func main() {
	SetCount(1)

	fmt.Println(GetCount())
}
