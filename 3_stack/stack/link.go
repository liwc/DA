package stack

import (
	"fmt"
)

// LinkStack ...
type LinkStack struct {
	Base *Node
	STop *Node
	Length int
}

// Push ...
func (ls *LinkStack) Push(data Item) {
	node := &Node{Data: data}
	if ls.STop == nil {
		ls.STop = node
		ls.Base = ls.STop
	} else {
		ls.STop.Next = node
		node.Prev = ls.STop
		ls.STop = ls.STop.Next
	}

	ls.Length++
}

// Pop ...
func (ls *LinkStack) Pop() Item {
	var res Item
	if ls.STop != nil {
		res = ls.STop.Data
		ls.STop = ls.STop.Prev
		if ls.STop != nil {
			ls.STop.Next = nil
		} else {
			ls.Base = ls.STop
		}

		ls.Length--
	}

	return res
}

// Top ...
func (ls *LinkStack) Top() Item {
	
	return nil
}

// Size ...
func (ls *LinkStack) Size() int {
	return ls.Length
}

// Empty ...
func (ls *LinkStack) Empty() bool {
	return ls.Length == 0
}

// Output ...
func (ls *LinkStack) Output() {
	p := ls.Base
	if p == nil {
		fmt.Println(p)
	}

	for p != nil {
		fmt.Printf("%v ", p.Data)
		p = p.Next
	}

	fmt.Println()
}

// LsInit ...
func LsInit() *LinkStack {
	ls := &LinkStack{}
	return ls
}

// LsTest ...
func LsTest() {
	ls := LsInit()
	fmt.Printf("Empty: %v\n", ls.Empty())

	for i := 0; i < 5; i++ {
		ls.Push(i)
	}
	ls.Output()

	fmt.Printf("Pop: %v\n", ls.Pop())
	ls.Output()
	fmt.Printf("Pop: %v\n", ls.Pop())
	ls.Output()
	fmt.Printf("Pop: %v\n", ls.Pop())
	ls.Output()
	fmt.Printf("Pop: %v\n", ls.Pop())
	ls.Output()
	fmt.Printf("Empty: %v\n", ls.Empty())
	fmt.Printf("Pop: %v\n", ls.Pop())
	ls.Output()
	fmt.Printf("Empty: %v\n", ls.Empty())
}