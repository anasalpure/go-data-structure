package main

import (
	"fmt"
	"strings"
)

/*
BST

	       50
	    /     \
	   30      70
	  /  \    /  \
	20   40  60   80
*/
func main() {
	fmt.Println("Hello, BST!")
	bst := Tree{}
	bst.insert(50)
	bst.insert(20)
	bst.insert(30)
	bst.insert(70)
	bst.insert(60)
	bst.insert(10)

	fmt.Println(bst.size)
	fmt.Println(bst)
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (node Node) String() string {
	return fmt.Sprintf("%d", node.value)
}

func (node *Node) insert(el int) {

	if el > node.value {
		// go right
		if node.right == nil {
			node.right = &Node{value: el}
		} else {
			node.right.insert(el)
		}
	} else {
		// go left
		if node.left == nil {
			node.left = &Node{value: el}
		} else {
			node.left.insert(el)
		}
	}
}

type Tree struct {
	head *Node
	size int
}

func (tree *Tree) insert(el int) {
	if tree.head == nil {
		newNode := Node{value: el}
		tree.head = &newNode
	} else {
		tree.head.insert(el)
	}
	tree.size++
}

func (tree Tree) printInOrder(node *Node, sb *strings.Builder) {
	if node == nil {
		return
	} else {
		tree.printInOrder(node.left, sb)
		sb.WriteString(fmt.Sprintf("%d ", node.value))
		tree.printInOrder(node.right, sb)
	}
}

func (tree Tree) String() string {
	sb := strings.Builder{}
	tree.printInOrder(tree.head, &sb)
	return sb.String()
}
