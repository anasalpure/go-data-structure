package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Stack!")
	stack := new(Stack)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	top, _ := stack.top()
	fmt.Println(top)
	fmt.Println("-------")

	for !stack.isEmpty() {
		val, _ := stack.pop()
		fmt.Println(val)
	}

}

type Node struct {
	value int
	next  *Node
}

func (node Node) String() string {
	return fmt.Sprintf("%d", node.value)
}

type Stack struct {
	head   *Node
	length int
}

func (stack *Stack) push(el int) {
	// stack.head = &Node{el, stack.head}

	newNode := Node{value: el}
	if stack.isEmpty() {
		stack.head = &newNode
	} else {
		newNode.next = stack.head
		stack.head = &newNode
	}
	stack.length++
}

func (stack *Stack) pop() (int, bool) {
	if !stack.isEmpty() {
		val := stack.head.value
		stack.head = stack.head.next
		stack.length--
		return val, true
	}
	return 0, false
}

func (stack *Stack) top() (int, bool) {
	if stack.head != nil {
		return stack.head.value, true
	}
	return 0, false
}

func (stack *Stack) isEmpty() bool {
	return stack.head == nil
}

func (stack *Stack) String() string {
	sb := strings.Builder{}
	for current := stack.head; current != nil; current = current.next {
		sb.WriteString(fmt.Sprintf("%d, ", current.value))
	}
	return sb.String()
}
