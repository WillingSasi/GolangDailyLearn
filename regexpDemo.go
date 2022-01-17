package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	buf := "abc azc z7c a7c aac 888 a9c  tac 43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8   "

	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	result1 := reg1.FindAllString(buf, -1)
	fmt.Println("result1 = ", result1)

	reg2 := regexp.MustCompile(`a\dc`)
	if reg2 == nil {
		fmt.Println("regexp err")
		return
	}
	result2 := reg2.FindAllString(buf, -1)
	fmt.Println("result2 = ", result2)

	reg3 := regexp.MustCompile(`a[0-9]c`)
	if reg3 == nil {
		fmt.Println("regexp err")
		return
	}
	result3 := reg3.FindAllString(buf, -1)
	fmt.Println("result3 = ", result3)

	reg4 := regexp.MustCompile(`\d+\.\d+`)
	if reg4 == nil {
		fmt.Println("regexp err")
		return
	}
	result4 := reg4.FindAllString(buf, -1)
	fmt.Println("result4 = ", result4)

	// 原生字符串
	buf1 := `<!DOCTYPE html>
		<html lang="zh-CN">
		<head>
		<title>C语言中文网 | Go语言入门教程</title>
		</head>
		<body>
		<div>Go语言简介</div>
		<div>Go语言基本语法
		Go语言变量的声明
		Go语言教程简明版
		</div>
		<div>Go语言容器</div>
		<div>Go语言函数</div>
		</body>
		</html>
		`
	reg5 := regexp.MustCompile(`<div>(?s:(.*?))<div>`)
	if reg5 == nil {
		fmt.Println("regexp err")
		return
	}
	result5 := reg5.FindAllString(buf1, -1)
	fmt.Println("result5 = ", result5)

	test()
}

func test() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("match found")
	}

	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str)
}
