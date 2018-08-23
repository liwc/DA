package queue

import (
	"fmt"
)

// OrderQueue ...
type OrderQueue struct {
	Front int
	Rear int
	Length int
	Buff [SIZE]Node
}

// Push ...
func (oq *OrderQueue) Push(data Item) bool {
	if oq.Full() {
		fmt.Println("queue full")
		return false
	}

	oq.Buff[oq.Rear] = Node{Data: data}
	oq.Rear = (oq.Rear + 1) % SIZE
	return true
}

// Pop ...
func (oq *OrderQueue) Pop() Item {
	if oq.Empty() {
		fmt.Println("queue empty")
		return nil
	}

	res := oq.Buff[oq.Front].Data
	oq.Front = (oq.Front + 1) % SIZE
	return res
}

// Top ...
func (oq *OrderQueue) Top() Item {
	return oq.Buff[oq.Front].Data
}

// Back ...
func (oq *OrderQueue) Back() Item {
	pos := (oq.Rear - 1 + SIZE) % SIZE 
	return oq.Buff[pos].Data
}

// Clear ...
func (oq *OrderQueue) Clear() {
	oq.Rear = oq.Front
}

// Empty ...
func (oq *OrderQueue) Empty() bool {
	if oq.Rear == oq.Front {
		return true
	}
	return false
}

// Full ...
func (oq *OrderQueue) Full() bool {
	if (oq.Rear + 1) % SIZE == oq.Front {
		return true
	}
	return false
}

// Size ...
func (oq *OrderQueue) Size() int {
	return (oq.Rear + SIZE - oq.Front) % SIZE
}

// Output ...
func (oq *OrderQueue) Output() {
	if oq.Empty() {
		fmt.Println("queue empty")
		return
	}

	pos := oq.Front
	for pos != oq.Rear {
		fmt.Printf("%v ", oq.Buff[pos].Data)
		pos = (pos + 1) % SIZE
	}

	fmt.Println()
}

// OqInit ...
func OqInit() *OrderQueue {
	return &OrderQueue{}
}

// OqTest ...
func OqTest() {
	fmt.Println("order queue test begin")
	oq := OqInit()
	oq.Output()
	fmt.Println(oq.Pop())

	oq.Push(0)
	oq.Push(1)
	oq.Push(2)
	fmt.Printf("len: %d\n", oq.Size())
	oq.Output()

	oq.Push(3)
	oq.Push(4)
	fmt.Printf("len: %d\n", oq.Size())
	oq.Output()

	fmt.Println(oq.Pop())
	fmt.Println(oq.Pop())
	fmt.Printf("len: %d\n", oq.Size())
	oq.Output()

	fmt.Println(oq.Pop())
	fmt.Println(oq.Pop())
	fmt.Println(oq.Pop())
	fmt.Printf("len: %d\n", oq.Size())
	oq.Output()

	fmt.Printf("order queue test end\n\n")
}