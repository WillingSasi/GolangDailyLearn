package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")

		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	<-ch

	fmt.Println("all done")
}
