package main

import (
	"fmt"
	"time"
)

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		// v[i] += u.Op(v[i])
	}

	c <- 1
}

const NCPU = 8

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	for i := 0; i < NCPU; i++ {
		<-c
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}
	time.Sleep(5 * time.Second)
}

func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}

	fmt.Printf("Thread %d, sum is : %d \n", index, sum)
}
