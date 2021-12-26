package main

import (
	"fmt"
	"math"
)

func main() {
	// var a int = 10

	// if a < 20 {
	// 	fmt.Println("a 小于 20")
	// }

	// fmt.Printf("a 的值为%d \n", a)

	// var aa int = 100
	// var bb int = 200
	// var ret int

	// ret = max(aa, bb)
	// fmt.Println(ret)

	// aaa, bbb := swap("Google", "Runoob")
	// fmt.Println(aaa, bbb)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}

func swap(x, y string) (string, string) {
	return y, x
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func testIf() {
	for true {
		fmt.Printf("无限循环\n")
	}
}

func test1() {
	var a int = 100

	if a < 20 {
		fmt.Println("a 小于 20")
	} else {
		fmt.Println("a 大于 20")
	}
}

func testSwitch() {
	var grade string = "s"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀")
	case grade == "B", grade == "C":
		fmt.Printf("良好")
	case grade == "D":
		fmt.Printf("及格")
	case grade == "F":
		fmt.Printf("不及格")
	default:
		fmt.Printf("差")
	}

	fmt.Printf("你的等级是 %s", grade)
}

func testTypeSwitch() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T", i)
	case int:
		fmt.Printf("x 是int")
	case float64:
		fmt.Printf("x 是float64型")
	case func(int) float64:
		fmt.Printf("x 是func(int) 型")
	case bool, string:
		fmt.Printf("x 是bool或string型")
	default:
		fmt.Printf("未定义型")
	}
}

func testFallthrough() {
	switch {
	case false:
		fmt.Printf("false")
		fallthrough
	case true:
		fmt.Printf("true")
		fallthrough
	case false:
		fmt.Printf("false")
		fallthrough
	case false:
		fmt.Printf("false")
		fallthrough
	case true:
		fmt.Printf("true")
	}
}
