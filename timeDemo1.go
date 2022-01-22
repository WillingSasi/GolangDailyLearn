package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	later1 := now.Add(-time.Hour)
	fmt.Println(later1)

	later2 := now.Sub(later1)
	fmt.Println(later2)

	fmt.Println(now.Equal(now))

	fmt.Println(now.Before(later))

	fmt.Println(now.After(later))
}
