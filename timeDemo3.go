package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	fmt.Println(now.Format("2006/01/02"))
}
