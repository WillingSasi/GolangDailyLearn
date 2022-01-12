package main

import (
	"fmt"

	"../model"
)

func main() {
	p := model.NewPerson("guanhua")
	p.SetAge(28)
	p.SetSal(10000)

	fmt.Println(p)
}
