package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	var ppptr ***int

	a = 3000

	ptr = &a

	pptr = &ptr

	ppptr = &pptr

	fmt.Printf("a = %d\n", a)
	fmt.Printf("*ptr = %d\n", *ptr)
	fmt.Printf("**pptr = %d\n", **pptr)
	fmt.Printf("***ppptr = %d\n", ***ppptr)
	fmt.Printf("ptr = %x\n", ptr)
	fmt.Printf("pptr = %x\n", pptr)
	fmt.Printf("ppptr = %x\n", ppptr)

	fmt.Printf("0 1 %b\n", 1024)
}
