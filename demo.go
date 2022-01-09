package main

import "fmt"

/*
duo hang zhu shi
*/

// dan hang zhu shi

var age = 1
var gc, gd int
var (
	ga int
	gb bool
)

const (
	Unknow = 0
	Female = 1
	Male   = 2
)

const (
	aa = iota
	bb
	cc
	dd = "ha"
	ee
	ff = 100
	gg
	hh = iota
	ii
)

func main() {
	fmt.Println("OKKKK" + "OKKKKK")
	fmt.Println("菜鸟教程")

	var stockcode int = 123
	var enddate string = "2021-12-1"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var a string
	fmt.Println(a)

	var b int
	fmt.Println(b)

	var c bool
	fmt.Println(c)

	var d error
	fmt.Println(d)

	e := "Runoob"
	fmt.Println(e)

	fmt.Println(&b)
	fmt.Println(&d)

	const f = 1
	fmt.Println(f)

	fmt.Println(hh)
	fmt.Println(ii)
}
