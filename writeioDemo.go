package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("你好go语言,入门教程")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后未使用的缓冲区为：%d\n", string(p), w.Available())
}
