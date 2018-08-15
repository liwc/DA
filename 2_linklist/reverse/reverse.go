package reverse

import (
	"fmt"
)

// Node ...
type Node struct {
	data int
	next *Node
}

// Reverse1 ...
func Reverse1(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	prev := head
	curr := prev.next
	prev.next = nil

	rear := curr.next
	for rear != nil {	
		curr.next = prev

		prev = curr
		curr = rear
		rear = rear.next
	}

	curr.next = prev
	head = curr
	
	fmt.Println()
	return head
}

// Reverse2 ...
func Reverse2(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	temp := &Node{}
	temp.next = head
	iter := head.next
	
	for iter != nil {
		head.next = iter.next
		iter.next = temp.next
		temp.next = iter

		iter = head.next
	}
	
	fmt.Println()
	return temp.next
}

// Reverse3 ...
func Reverse3(curr *Node) (*Node, *Node) {
	if curr == nil || curr.next == nil {
		return curr, curr
	}

	rear, tail := Reverse3(curr.next)
	curr.next = nil
	rear.next = curr

	return curr, tail
}

// Test ...
func Test() {
	head := &Node{0, nil}
	tail := head
	for i := 2; i <= 5; i++ {
		node := &Node{data: i}
		tail.next = node
		tail = node
	}

	fmt.Println("reverse before:")
	for iter := head; iter != nil; iter = iter.next {
		fmt.Printf("%d ", iter.data)
	}

	_, head = Reverse3(head)

	fmt.Println("\nreverse after:")
	for iter := head; iter != nil; iter = iter.next {
		fmt.Printf("%d ", iter.data)
	}
}