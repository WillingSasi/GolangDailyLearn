package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)
	var l sync.Mutex

	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解了，你们去抢吧")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("goroutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 我也解锁了")
		ch <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-ch
	}
}
