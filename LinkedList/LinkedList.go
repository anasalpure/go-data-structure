package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Println("Hello, Linked List!")
	linkedList := new(LinkedList)
	linkedList.add(1)
	linkedList.add(2)
	linkedList.add(3)
	linkedList.add(4)
	linkedList.remove(1)
	linkedList.remove(3)
	fmt.Print(linkedList)
	fmt.Print(linkedList.length)
}

type Node struct{
	value int
	next *Node
}
func (node Node) String () string{
	return fmt.Sprintf("%d", node.value)
}

type LinkedList struct {
	head *Node
	length int 
}

func (list *LinkedList) add(el int){
	newNode := &Node{value: el}

	if  list.head == nil {
		list.head = newNode
	} else {
		var current *Node = list.head
		for ; current.next != nil; current = current.next {}
		current.next = newNode
	}
	list.length++
}

func (list *LinkedList) remove(el int) bool {
	removed := false;
	var previous *Node
	// remove first node
	if list.head != nil && list.head.value == el {
		list.head = list.head.next
		list.length --
		return true

	}

	for current := list.head; current.next !=nil; current= current.next {
		if current.value == el {
			previous.next = current.next
			list.length --
			removed = true
		}
		previous = current;
	}
	return removed
}

func (list *LinkedList) String() string {
	// %v the value in a default format
	// %+v the plus flag adds field names
	// %#v a Go-syntax representation of the value
	// fmt.Printf("%#v",list)
	// fmt.Println()


	sb := strings.Builder{}
	for current:=list.head ; current != nil; current = current.next {
		sb.WriteString(fmt.Sprintf("%d, ",current.value))
	}
	return sb.String()
}