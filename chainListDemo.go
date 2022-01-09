package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
}

func main() {
	var head = new(Node)
	head.Data = 0
	var tail *Node
	tail = head

	for i := 1; i < 10; i++ {
		var node = Node{Data: i}
		(*tail).Next = &node
		tail = &node
	}

	ShowNode(head)
}

func main2() {
	var head = new(Node)
	head.Data = 0
	var tail *Node
	tail = head

	for i := 1; i < 10; i++ {
		var node = Node{Data: i}
		node.Next = tail
		tail = &node
	}

	ShowNode(tail)
}

func main1() {
	var head = new(Node)
	head.Data = 1
	var node1 = new(Node)
	node1.Data = 2

	head.Next = node1
	var node2 = new(Node)
	node2.Data = 3

	node1.Next = node2

	ShowNode(head)
}
