package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ())

	fmt.Println(os.Hostname())
}
