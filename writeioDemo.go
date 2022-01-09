package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	// p := []byte("你好go语言,入门教程")

	// fmt.Println("写入前未使用的缓冲区为：", w.Available())
	// w.Write(p)
	// fmt.Printf("写入%q后未使用的缓冲区为：%d\n", string(p), w.Available())

	// fmt.Println("写入前未使用的缓冲区为：", w.Buffered())
	// w.Write(p)
	// fmt.Printf("写入%q后未使用的缓冲区为：%d\n", string(p), w.Buffered())
	// w.Flush()
	// fmt.Println("执行Flush后，写入的字节数为：", w.Buffered())

	// w.Write(p)
	// fmt.Printf("未执行Flush，缓冲区输出%q:\n", string(wr.Bytes()))
	// w.Flush()
	// fmt.Printf("执行Flush，缓冲区输出%q:\n", string(wr.Bytes()))

	// n, err := w.Write(p)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), n, err)

	// var c byte = 'G'
	// err := w.WriteByte(c)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), err)

	// var r rune = 'G'
	// size, err := w.WriteRune(r)
	// w.Flush()
	// fmt.Println(string(wr.Bytes()), size, err)

	s := "C语言中文网学习笔记"
	size, err := w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)
}
