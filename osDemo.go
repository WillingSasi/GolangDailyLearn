package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(os.Environ())

	fmt.Println(os.Hostname())

	fmt.Println(os.Geteuid())

	fmt.Println(os.Getpid())
}
