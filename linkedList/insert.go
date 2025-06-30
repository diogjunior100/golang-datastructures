package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

func (list *LinkedList) Display() {
	currentNode := list.head

	if currentNode == nil {
		fmt.Printf("Linked list is empty")
	}

	fmt.Printf("Linked list:\n")
	for currentNode != nil {
		if currentNode.next != nil {
			fmt.Printf("%d->", currentNode.data)
		} else {
			fmt.Printf("%d", currentNode.data)
		}
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Display()

}
