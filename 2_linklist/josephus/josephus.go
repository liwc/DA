package josephus

import (
	"fmt"
)

// Node ...
type Node struct {
	data int
	next *Node
}

// Josephus ...
func Josephus(n, k int) int {
	head := &Node{1, nil}
	tail := head
	for i := 2; i <= n; i++ {
		node := &Node{i, nil}
		tail.next = node
		tail = node
	}
	tail.next = head

	iter := tail
	count := n
	for count > 1 {
		count--

		for i := 1; i < k; i++ {
			iter = iter.next
		}
		fmt.Printf("%d ", iter.next.data)
		iter.next = iter.next.next
	}

	fmt.Println()
	return iter.data
}