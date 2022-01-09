package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("你好go语言,入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	// var buf [128]byte
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)

	// c, err := r.ReadByte()
	// fmt.Println(string(c), err)

	// var delim byte = ','
	// line, err := r.ReadBytes(delim)
	// fmt.Println(string(line), err)

	// line, prefix, err := r.ReadLine()
	// fmt.Println(string(line), prefix, err)

	// ch, size, err := r.ReadRune()
	// fmt.Println(string(ch), size, err)

	// var delim byte = ','
	// line, err := r.ReadSlice(delim)
	// fmt.Println(string(line), err)
	// line, err = r.ReadSlice(delim)
	// fmt.Println(string(line), err)
	// line, err = r.ReadSlice(delim)
	// fmt.Println(string(line), err)

	// var delim byte = ','
	// line, err := r.ReadString(delim)
	// fmt.Println(string(line), err)

	// var buf [14]byte
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)
	// rn := r.Buffered()
	// fmt.Println(rn)
	// n, err = r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)
	// rn = r.Buffered()
	// fmt.Println(rn)

	bl, err := r.Peek(8)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(14)
	fmt.Println(string(bl), err)
	//1中文=3字节，所以不是3的整数倍，读的字符显示不完整
	bl, err = r.Peek(20)
	fmt.Println(string(bl), err)
	bl, err = r.Peek(21)
	fmt.Println(string(bl), err)
}
