package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)

	signal.Stop(c)
	s := <-c
	fmt.Println("Got signal: ", s)
}
