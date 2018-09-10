package tree

import (
	"fmt"
)

// Item ...
type Item interface {}

// Node ...
type Node struct {
	Data Item
	Left *Node
	Right *Node
	Parent *Node

	Reserve1 Item
	Reserve2 Item
	Reserve3 Item
}

func preorderTraverseImpl(node *Node) {
	if node != nil {
		fmt.Printf("%v ", node.Data)
		preorderTraverseImpl(node.Left)
		preorderTraverseImpl(node.Right)
	}
}

func inorderTraverseImpl(node *Node) {
	if node != nil {
		inorderTraverseImpl(node.Left)
		fmt.Printf("%v ", node.Data)
		inorderTraverseImpl(node.Right)
	}
}

func postorderTraverseImpl(node *Node) {
	if node != nil {
		postorderTraverseImpl(node.Left)
		postorderTraverseImpl(node.Right)
		fmt.Printf("%v ", node.Data)
	}
}