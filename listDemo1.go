package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	// var l2 list.List

	l.PushBack("first")
	l.PushFront(67)
	l.PushBack("cannon")

	element := l.PushBack("fist")

	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
