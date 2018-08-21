package stack

import (
	"fmt"
)

// STACKSIZE ...
const STACKSIZE int = 5
// INCREMENT ...
const INCREMENT int = 5

// OrderStack ...
type OrderStack struct {
	Base int
	Stop int
	Length int
	Buff []Node
	Size int
}

// Push ...
func (os *OrderStack) Push(data Item) {
	if os.Length >= os.Size {
		tmpBuf := make([]Node, os.Size)
		copy(tmpBuf, os.Buff)

		os.Buff = make([]Node, os.Size + INCREMENT)
		copy(os.Buff, tmpBuf)
		os.Size += INCREMENT
	}

	os.Buff[os.Stop] = Node{Data: data}
	os.Stop++
	os.Length++
}

// Pop ...
func (os *OrderStack) Pop() Item {
	if os.Stop > 0 {
		os.Stop--
		os.Length--
		return os.Buff[os.Stop].Data
	}

	return nil
}

// Top ...
func (os *OrderStack) Top() Item {
	if os.Stop > 0 {
		return os.Buff[os.Stop - 1].Data
	}
	return nil
}

// Capacity ...
func (os *OrderStack) Capacity() int {
	return os.Size
}

// Empty ...
func (os *OrderStack) Empty() bool {
	return os.Stop == 0
}

// Output ...
func (os *OrderStack) Output() {
	for i := 0; i < os.Stop; i++ {
		fmt.Printf("%v ", os.Buff[i].Data)
	}

	fmt.Println()
}

// OsInit ...
func OsInit() *OrderStack {
	os := &OrderStack{
		Base: 0,
		Stop: 0,
		Length: 0,
		Buff: make([]Node, STACKSIZE),
		Size: STACKSIZE,
	}

	return os
}

// OsTest ...
func OsTest() {
	os := OsInit()
	fmt.Printf("Empty: %v\n", os.Empty())

	for i := 0; i < 5; i++ {
		os.Push(i)
	}
	os.Output()

	for i := 5; i < 8; i++ {
		os.Push(i)
	}
	os.Output()

	fmt.Printf("Pop: %v\n", os.Pop())
	os.Output()

	fmt.Printf("Top: %v\n", os.Top())
	fmt.Printf("Empty: %v\n", os.Empty())
}