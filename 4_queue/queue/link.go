package queue

import (
	"fmt"
)

// LinkQueue ...
type LinkQueue struct {
	Front *Node
	Rear *Node
	Length int
}

// Push ...
func (lq *LinkQueue) Push(data Item) {
	node := &Node{Data: data}
	lq.Rear.Next = node
	lq.Rear = lq.Rear.Next
	
	lq.Length++
}

// Pop ...
func (lq *LinkQueue) Pop() Item {
	if lq.Empty() {
		return nil
	}

	lq.Length--
	lq.Front = lq.Front.Next
	return lq.Front.Data
}

// Top ...
func (lq *LinkQueue) Top() Item {
	if lq.Empty() {
		return nil
	}
	return lq.Front.Next.Data
}

// Back ...
func (lq *LinkQueue) Back() Item {
	if lq.Empty() {
		return nil
	}
	return lq.Rear.Data
}

// Clear ...
func (lq *LinkQueue) Clear() {
	lq.Rear = lq.Front
}

// Empty ...
func (lq *LinkQueue) Empty() bool {
	if lq.Front == lq.Rear {
		fmt.Println("queue empty")
		return true
	}
	return false
}

// Size ...
func (lq *LinkQueue) Size() int {
	return lq.Length
}

// Output ...
func (lq *LinkQueue) Output() {
	if lq.Empty() {
		return
	}

	iter := lq.Front.Next
	for iter != nil {
		fmt.Printf("%v ", iter.Data)
		iter = iter.Next
	}

	fmt.Println()
}

// LqInit ...
func LqInit() *LinkQueue {
	node := &Node{}
	return &LinkQueue{Front: node, Rear: node}
}

// LqTest ...
func LqTest() {
	fmt.Println("link queue test begin")

	lq := LqInit()
	lq.Output()
	fmt.Printf("top: %v\n", lq.Top())
	fmt.Printf("back: %v\n", lq.Back())

	lq.Push(0)
	lq.Push(1)
	fmt.Printf("len: %d\n", lq.Size())
	lq.Output()
	fmt.Printf("top: %v\n", lq.Top())
	fmt.Printf("back: %v\n", lq.Back())

	fmt.Printf("pop: %v\n", lq.Pop())
	fmt.Printf("pop: %v\n", lq.Pop())
	lq.Output()
	fmt.Printf("top: %v\n", lq.Top())
	fmt.Printf("back: %v\n", lq.Back())

	fmt.Printf("link queue test end\n\n")
}